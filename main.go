package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type SiteList struct {
	Url      string   `json:"url"`
	Sitename []string `json:"sitename"`
}

func main() {
	var cmd *exec.Cmd
	var UsingOSCmd string
	var err error

	// フラグ定義
	urlFlag := flag.String("url", "", "開きたいWebサイトの(https://から始まる)URLを入力してください。")
	siteFlag := flag.String("site", "", "(website-cli-setting.jsonで登録した)開きたいサイト名を入力してください。")
	searchFlag := flag.String("search", "", "検索したい用語を入力してください。スペースを入れて検索したい場合は、クオーテーションで囲むか単語+でつないでください。")
	flag.Parse()

	// 実行ファイルのパスを取得
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("os.Executable err:", err)
	}

	// 実行ファイルのディレクトリのパスに変換
	exeDir := filepath.Dir(exePath)

	// ディレクトリのパスに読み込む設定ファイルの名前を追加
	settingfilePath := filepath.Join(exeDir, "website-cli-setting.json")

	// 設定ファイル読み込み
	f, err := os.ReadFile(settingfilePath)
	if err != nil {
		log.Fatal("os.ReadFile err:", err)
	}

	// json読み込み
	var sl []SiteList
	err = json.Unmarshal(f, &sl)
	if err != nil {
		log.Fatal("json.Unmarshal err:", err)
	}

	// 使用しているOS確認
	UsingOSCmd, err = OSCheck()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *siteFlag != "" {
		// サイト名選択
		cmd, err = Site(siteFlag, UsingOSCmd, sl)
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			os.Exit(1)
		}

		// websiteを開く
		err = OpenWebSite(cmd)
		if err != nil {
			fmt.Println("Webサイトを開く際にエラーが発生しました:", err)
			os.Exit(1)
		}
	}

	if *urlFlag != "" {
		// URL指定
		cmd, err = Url(urlFlag, UsingOSCmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// websiteを開く
		err = OpenWebSite(cmd)
		if err != nil {
			fmt.Println("Webサイトを開く際にエラーが発生しました:", err)
			os.Exit(1)
		}
	}

	if *searchFlag != "" {
		// 検索ワード指定
		cmd, err = Search(searchFlag, UsingOSCmd)
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			os.Exit(1)
		}

		// websiteを開く
		err = OpenWebSite(cmd)
		if err != nil {
			fmt.Println("Webサイトを開く際にエラーが発生しました:", err)
			os.Exit(1)
		}
	}
}

func OSCheck() (string, error) {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows使用中")
		return "start", nil
	case "darwin":
		fmt.Println("Mac OSX使用中")
		return "open", nil
	case "linux":
		fmt.Println("Linux使用中")
		return "xgd-open", nil
	default:
		return "", errors.New("対応外のOSです")
	}
}

func Site(siteFlag *string, UsingOSCmd string, sitelist []SiteList) (*exec.Cmd, error) {
	// サイト名選択
	for _, site := range sitelist {
		for _, sn := range site.Sitename {
			if *siteFlag == sn {
				cmd := exec.Command(UsingOSCmd, site.Url)
				return cmd, nil
			}
		}
	}
	return nil, errors.New("登録されていないサイト名です")
}

func Url(urlFlag *string, UsingOSCmd string) (*exec.Cmd, error) {
	s_urlflag := *urlFlag
	// URLが正しく入力されているか
	if len(s_urlflag) < 7 {
		return nil, errors.New("URLのプロトコルに誤りがあります")
	} else if s_urlflag[:8] != "https://" && s_urlflag[:7] != "http://" {
		return nil, errors.New("URLのプロトコルに誤りがあります")
	}
	cmd := exec.Command(UsingOSCmd, *urlFlag)
	return cmd, nil
}

func Search(searchFlag *string, UsingOSCmd string) (*exec.Cmd, error) {
	// URL作成
	url := "https://google.com"
	path := "/search"
	params := "?q=" + *searchFlag
	cmd := exec.Command(UsingOSCmd, url+path+params)
	return cmd, nil
}

func OpenWebSite(cmd *exec.Cmd) error {
	// コマンドを実行してWebサイトを開く
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
