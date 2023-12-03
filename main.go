package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd
	var usingoscmd string

	// フラグ定義
	urlFlag := flag.String("site", "", "開きたいWebサイトの(https://から始まる)URLやサイト名(Twitter,ニコニコ動画,YouTube,GitHub,Classroom)を入力してください")
	flag.Parse()

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows使用中")
		usingoscmd = "start"
	case "darwin":
		fmt.Println("Mac OSX使用中")
		usingoscmd = "open"
	case "linux":
		fmt.Println("Linux使用中")
		usingoscmd = "xgd-open"
	default:
		fmt.Println("対応外のOSです")
		os.Exit(1)
	}

	switch {
	case *urlFlag == "":
		fmt.Println("何も指定されていません。")
		flag.PrintDefaults()
		os.Exit(1)
	case *urlFlag == "twitter" || *urlFlag == "Twitter" || *urlFlag == "x" || *urlFlag == "X" || *urlFlag == "ツイッター" || *urlFlag == "ついったー":
		cmd = exec.Command(usingoscmd, "https://twitter.com")
	case *urlFlag == "niconico" || *urlFlag == "ニコニコ" || *urlFlag == "ニコニコ動画" || *urlFlag == "にこにこ" || *urlFlag == "にこにこどうが" || *urlFlag == "ニコ動" || *urlFlag == "nikoniko":
		cmd = exec.Command(usingoscmd, "https://www.nicovideo.jp")
	case *urlFlag == "youtube" || *urlFlag == "YouTube" || *urlFlag == "ようつべ" || *urlFlag == "ユーチューブ":
		cmd = exec.Command(usingoscmd, "https://www.youtube.com")
	case *urlFlag == "github" || *urlFlag == "GitHub" || *urlFlag == "ギットハブ" || *urlFlag == "ぎっとはぶ":
		cmd = exec.Command(usingoscmd, "https://github.com")
	case *urlFlag == "classroom" || *urlFlag == "Classroom" || *urlFlag == "クラスルーム" || *urlFlag == "くらするーむ":
		cmd = exec.Command(usingoscmd, "https://classroom.google.com")
	default:
		s_urlflag := *urlFlag
		if len(s_urlflag) < 7 {
			fmt.Println("URLのプロトコルに誤りがあります")
			os.Exit(1)
		} else if s_urlflag[:8] != "https://" && s_urlflag[:7] != "http://" {
			fmt.Println("URLのプロトコルに誤りがあります。")
			os.Exit(1)
		}
		cmd = exec.Command(usingoscmd, *urlFlag)
	}

	// コマンドを実行してWebサイトを開く
	if err := cmd.Start(); err != nil {
		fmt.Println("Webサイトを開く際にエラーが発生しました:", err)
		os.Exit(1)
	}
}
