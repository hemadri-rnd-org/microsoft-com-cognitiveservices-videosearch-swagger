package main

import (
	"github.com/video-search-client/mcp-server/config"
	"github.com/video-search-client/mcp-server/models"
	tools_videodetailsearch "github.com/video-search-client/mcp-server/tools/videodetailsearch"
	tools_videosearch "github.com/video-search-client/mcp-server/tools/videosearch"
	tools_videotrendingsearch "github.com/video-search-client/mcp-server/tools/videotrendingsearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_videodetailsearch.CreateVideos_detailsTool(cfg),
		tools_videosearch.CreateVideos_searchTool(cfg),
		tools_videotrendingsearch.CreateVideos_trendingTool(cfg),
	}
}
