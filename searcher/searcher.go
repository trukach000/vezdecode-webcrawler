package searcher

import (
	"log"
	"vezdecode-webcrawler/searcher/entities"
	"vezdecode-webcrawler/utils"
)

type Searcher struct {
	EntitiesToSearch []entities.SearchEntity
}

func NewSearcher(entities ...entities.SearchEntity) *Searcher {
	return &Searcher{
		EntitiesToSearch: entities,
	}
}

func (s *Searcher) SearchAll(data string) map[string][]string {
	result := make(map[string][]string)

	for _, es := range s.EntitiesToSearch {
		res, err := es.Search(data)
		if err != nil {
			log.Printf("Error during search entity (%s): %s", es.Title(), err)
			continue
		}

		uniqRes := utils.Unique(res)

		result[es.Title()] = uniqRes
	}

	return result
}
