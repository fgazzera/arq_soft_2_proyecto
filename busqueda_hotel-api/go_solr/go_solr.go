package go_solr

import (
	"log"

	"github.com/vanng822/go-solr/solr"
)

var SolrClient *solr.SolrInterface

func InitSolrClient(solrURL string) {
	solrClient, err := solr.NewSolrInterface(solrURL, "busqueda_hotel-core")
	if err != nil {
		log.Fatalf("Error al conectar a Solr: %s", err)
	}
	SolrClient = solrClient
}

func GetSolrClient() *solr.SolrInterface {
	return SolrClient
}
