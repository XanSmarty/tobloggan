package integration

import (
	"io/fs"
	"sync/atomic"
	"time"

	"tobloggan/code/contracts"
	"tobloggan/code/stations"

	"github.com/mdwhatcott/pipelines"
)

type Config struct {
	Clock             contracts.Clock
	Logger            contracts.Logger
	MarkdownConverter stations.Markdown
	FileSystemReader  fs.FS
	FileSystemWriter  contracts.FSWriter
	TargetDirectory   string
	ArticleTemplate   string
	ListingTemplate   string
	BaseURL           string
}

func GenerateBlog(config Config) bool {
	started := config.Clock()
	defer func() { config.Logger.Printf("finished in %s", time.Since(started)) }()

	var (
		failure = new(atomic.Bool)
		input   = make(chan any, 1)

		scanner   = pipelines.Station(nil) // stations.NewSourceScanner(config.FileSystemReader)
		reader    = pipelines.Station(nil) // stations.NewSourceReader(config.FileSystemReader)
		parser    = pipelines.Station(nil) // stations.NewArticleParser()
		validator = pipelines.Station(nil) // stations.NewArticleValidator() // OPTIONAL?
		drafts    = pipelines.Station(nil) // stations.NewDraftRemoval() // OPTIONAL
		futures   = pipelines.Station(nil) // stations.NewFutureRemoval(started)
		markdown  = pipelines.Station(nil) // stations.NewMarkdownConverter(config.MarkdownConverter)
		listing   = pipelines.Station(nil) // stations.NewListingRenderer(config.ListingTemplate)
		renderer  = pipelines.Station(nil) // stations.NewArticleRenderer(config.ArticleTemplate)
		baseURL   = stations.NewBaseURLRewriter(config.BaseURL)
		writer    = stations.NewPageWriter(config.TargetDirectory, config.FileSystemWriter)
		reporter  = stations.NewReporter(config.Logger, failure)

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
			pipelines.Options.StationSingleton(validator),
			pipelines.Options.StationSingleton(drafts),
			pipelines.Options.StationSingleton(futures),
			pipelines.Options.StationSingleton(markdown),
			pipelines.Options.StationSingleton(listing),
			pipelines.Options.StationSingleton(renderer),
			pipelines.Options.StationSingleton(baseURL),
			pipelines.Options.StationSingleton(writer), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(reporter),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failure.Load()
}
