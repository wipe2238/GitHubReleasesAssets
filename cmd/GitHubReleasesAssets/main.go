package main

import (
	"fmt"
	"strconv"
)

func main() {
	JSON := GetJSON(fmt.Sprintf("https://api.github.com/repos/%s/releases", CmdLine.Repository))

	var (
		ReleaseLen   int = 0
		AssetLen     int = 0
		DownloadsLen int = 0
	)

	for _, Release := range JSON.GetArray() {
		if Len := len(Release.GetStringBytes("name")); Len > ReleaseLen {
			// fmt.Printf("ReleaseLen %d -> %d\n", ReleaseLen, Len)
			ReleaseLen = Len
		}
		AssetArray := Release.GetArray("assets")
		if AssetArray == nil || len(AssetArray) < 1 {
			continue
		}
		for _, Asset := range AssetArray {
			if Len := len(Asset.GetStringBytes("name")); Len > AssetLen {
				// fmt.Printf("AssetLen %d -> %d\n", AssetLen, Len)
				AssetLen = Len
			}

			if Len := len(strconv.Itoa(Asset.GetInt("download_count"))); Len > DownloadsLen {
				// fmt.Printf("DownloadsLen %d -> %d\n", DownloadsLen, Len)
				DownloadsLen = Len
			}
		}
	}

	var Format = fmt.Sprintf("%%-%ds  %%-%ds  %%%dd  %%s\n", ReleaseLen, AssetLen, DownloadsLen)
	// fmt.Println(Format)

	for _, Release := range JSON.GetArray() {
		AssetArray := Release.GetArray("assets")
		if AssetArray == nil || len(AssetArray) < 1 {
			continue
		}

		var ReleaseName string = fmt.Sprintf("%s", Release.GetStringBytes("name"))

		for _, Asset := range AssetArray {
			fmt.Printf(Format,
				ReleaseName,
				Asset.GetStringBytes("name"),
				Asset.GetUint("download_count"),
				Asset.GetStringBytes("browser_download_url"))

			// when release have multiple assets, print its name only once
			ReleaseName = ""
		}
	}
}
