package database

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
	"time"
)

const ES_ACGER_PAIR = "acger_pair"

var AcgerPairBody = `
{
  "aliases": {
    "acger_pair": {}
  },
  "settings": {
    "number_of_replicas": 0,
    "number_of_shards": 5,
    "analysis": {
      "analyzer": {
        "ik_synonym": {
          "tokenizer": "ik_max_word",
          "filter": [
            "remote_synonym",
            "lowercase"
          ]
        }
      },
      "filter": {
        "remote_synonym": {
          "type": "dynamic_synonym",
          "synonyms_path": "http://source.acger-pair.com/es/tongyici/dict.txt",
          "interval": 30
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "uid": {
        "type": "integer"
      },
      "boost": {
        "type": "integer"
      },
      "start": {
        "type": "boolean"
      },
      "skill": {
        "type": "text",
        "analyzer": "ik_synonym"
      },
      "skill_need": {
        "type": "text",
        "analyzer": "ik_synonym"
      },
      "created_at": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis||yyyy-MM-dd HH:mm:ss"
      },
      "updated_at": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis||yyyy-MM-dd HH:mm:ss"
      },
      "deleted_at": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis||yyyy-MM-dd HH:mm:ss"
      }
    }
  }
}
`

type Shards struct {
	Total int `json:"total"`
	Successful int `json:"successful"`
	Skipped int `json:"skipped"`
	Failed int `json:"failed"`
}
type Total struct {
	Value int `json:"value"`
	Relation string `json:"relation"`
}
type Source struct {
	Skill string `json:"skill"`
	SkillNeed string `json:"skill_need"`
	UID int `json:"uid"`
	Boost int `json:"boost"`
	Star int `json:"star"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}
type Highlight struct {
	Skill []string `json:"skill"`
	SkillNeed []string `json:"skill_need"`
}
type Hits struct {
	Index string `json:"_index"`
	Type string `json:"_type"`
	ID string `json:"_id"`
	Score float64 `json:"_score"`
	Source Source `json:"_source"`
	Highlight Highlight `json:"highlight"`
}
type Hit struct {
	Total Total `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits []Hits `json:"hits"`
}

type EsSearchPairResult struct {
	Took int `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards Shards `json:"_shards"`
	Hits Hit `json:"hits"`
}

func InitIndex(client *es.Client, index string, body string) {
	rsp, err := client.Indices.Get([]string{index})

	if err != nil {
		logx.Error("init es index fail")
		logx.Error(err.Error())
		return
	}

	if rsp.StatusCode == http.StatusNotFound {
		_, e := client.Indices.Create(index, client.Indices.Create.WithBody(strings.NewReader(body)))
		if e != nil {
			logx.Error("init es index fail.")
			logx.Error(e.Error())
			return
		}
	}
}
