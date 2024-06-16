package main

import (
	"fmt"
	"strconv"
)

func main() {
	JSON := GetJSON(fmt.Sprintf("https://api.github.com/repos/%s/releases", CmdLine.Repository))

	var (
		ReleaseText   string = "Release"
		AssetText     string = "Asset"
		DownloadsText string = "D/L"
		UrlText       string = "URL"

		ReleaseLen   int = len(ReleaseText)
		AssetLen     int = len(AssetText)
		DownloadsLen int = len(DownloadsText)

		ReleaseFound bool = false
		AssetFound   bool = false
		AssetFirst   bool = true
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

			if Len := len(strconv.FormatUint(uint64(Asset.GetUint("download_count")), 10)); Len > DownloadsLen {
				// fmt.Printf("DownloadsLen %d -> %d\n", DownloadsLen, Len)
				DownloadsLen = Len
			}
		}
	}

	var Format = fmt.Sprintf("%%-%ds  %%-%ds  %%%ds  %%s\n", ReleaseLen, AssetLen, DownloadsLen)
	// fmt.Println(Format)

	for _, Release := range JSON.GetArray() {
		ReleaseFound = true

		AssetArray := Release.GetArray("assets")
		if AssetArray == nil || len(AssetArray) < 1 {
			continue
		}

		var ReleaseName string = string(Release.GetStringBytes("name"))

		for _, Asset := range AssetArray {
			AssetFound = true

			if AssetFirst {
				AssetFirst = false
				fmt.Printf(Format,
					ReleaseText,
					AssetText,
					DownloadsText,
					UrlText)
			}

			fmt.Printf(Format,
				ReleaseName,
				Asset.GetStringBytes("name"),
				strconv.FormatUint(uint64(Asset.GetUint("download_count")), 10),
				Asset.GetStringBytes("browser_download_url"))

			// when release have multiple assets, print its name only once
			ReleaseName = ""
		}
	}

	if !ReleaseFound {
		fmt.Print("Releases not found")
	} else if !AssetFound {
		fmt.Print("Releases assets not found")
	}
}
