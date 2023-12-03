package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd
	var UsingOSCmd string

	// フラグ定義
	urlFlag := flag.String("url", "", "開きたいWebサイトの(https://から始まる)URLを入力してください")
	siteFlag := flag.String("site", "", "開きたいサイト名(Twitter,ニコニコ動画,YouTube,GitHub,Classroom)を入力してください")
	flag.Parse()

	UsingOSCmd, err := OSCheck()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *siteFlag != "" {
		// サイト名選択
		switch {
		case *siteFlag == "twitter" || *siteFlag == "Twitter" || *siteFlag == "x" || *siteFlag == "X" || *siteFlag == "ツイッター" || *siteFlag == "ついったー":
			cmd = exec.Command(UsingOSCmd, "https://twitter.com")
		case *siteFlag == "niconico" || *siteFlag == "ニコニコ" || *siteFlag == "ニコニコ動画" || *siteFlag == "にこにこ" || *siteFlag == "にこにこどうが" || *siteFlag == "ニコ動" || *siteFlag == "nikoniko":
			cmd = exec.Command(UsingOSCmd, "https://www.nicovideo.jp")
		case *siteFlag == "youtube" || *siteFlag == "YouTube" || *siteFlag == "ようつべ" || *siteFlag == "ユーチューブ":
			cmd = exec.Command(UsingOSCmd, "https://www.youtube.com")
		case *siteFlag == "github" || *siteFlag == "GitHub" || *siteFlag == "ギットハブ" || *siteFlag == "ぎっとはぶ":
			cmd = exec.Command(UsingOSCmd, "https://github.com")
		case *siteFlag == "classroom" || *siteFlag == "Classroom" || *siteFlag == "クラスルーム" || *urlFlag == "くらするーむ":
			cmd = exec.Command(UsingOSCmd, "https://classroom.google.com")
		default:
			fmt.Println("登録されていないサイト名です")
			flag.PrintDefaults()
			os.Exit(1)
		}

		// websiteを開く
		err := OpenWebSite(cmd)
		if err != nil {
			fmt.Println("Webサイトを開く際にエラーが発生しました:", err)
			os.Exit(1)
		}
	}

	if *urlFlag != "" {
		s_urlflag := *urlFlag
		// URLが正しく入力されているか
		if len(s_urlflag) < 7 {
			fmt.Println("URLのプロトコルに誤りがあります")
			os.Exit(1)
		} else if s_urlflag[:8] != "https://" && s_urlflag[:7] != "http://" {
			fmt.Println("URLのプロトコルに誤りがあります。")
			os.Exit(1)
		}
		cmd = exec.Command(UsingOSCmd, *urlFlag)

		// websiteを開く
		err := OpenWebSite(cmd)
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

func OpenWebSite(cmd *exec.Cmd) error {
	// コマンドを実行してWebサイトを開く
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
