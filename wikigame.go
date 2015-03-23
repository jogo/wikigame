package main

import (
	"flag"
	"fmt"
	"github.com/jogo/wikigame/graph"
	"github.com/jogo/wikigame/xml"
	"log"
	"os"
	"time"
)

var filename = "pages.bolt"

func importXML(limit *int) {
	//create graph
	g := graph.NewGraph(filename)
	defer g.Close()

	//populate graph from enwiki*.xml.bz2
	// bzcat enwiki*.xml.bz2 | ./me
	p, err := xml.NewParser(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; err == nil && i < *limit; i++ {
		var page *xml.Page
		page, err = p.Next()
		if err == nil {
			links := xml.FindLinks(page.Text)
			if i%1000 == 0 {
				//print every 1000th input
				fmt.Printf("%d, %s, %d\n", i, page.Title, len(links))
			}

			g.Put(string(page.Title), links)
		}
		g.Flush()
	}
	return
}

func timePath(g *graph.Graph, source, dest string) {
	start := time.Now()
	path := g.Path(source, dest)
	elapsed := time.Since(start)
	fmt.Printf("%s [%s]\n", path, elapsed)
}

func test(g *graph.Graph) {
	fmt.Println("Belgium => Economy of Belgium: ", g.Linked("Belgium", "Economy of Belgium"))
	fmt.Println("Belgium => beer: ", g.Linked("Belgium", "beer"))
	fmt.Println("Beer => stout: ", g.Linked("Beer", "stout"))
	timePath(g, "Belgium", "stout")
	timePath(g, "Belgium", "Baltic region")
	timePath(g, "Baltic region", "Belgium")
	timePath(g, "Belgium", "Budweiser (Anheuser-Busch)")
	timePath(g, "Budweiser (Anheuser-Busch)", "Belgium")
	timePath(g, "ozone depletion", "mexican war")
}

func main() {
	importGraph := flag.Bool("import", false, "Import new graph instead of using saved one")
	limit := flag.Int("limit", 500000, "Stop after parsing limit number of pages. For testing purposes")

	flag.Parse()
	if *importGraph == true {
		fmt.Println("Usage: bzcat enwiki*.xml.bz2 | ./", os.Args[0])
		importXML(limit)
	}

	g := graph.NewGraph(filename)
	defer g.Close()
	test(g)
}
