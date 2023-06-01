package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	//info, err := os.Lstat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf(" ファイル名: %v\n", info.Name())
	fmt.Printf(" サイズ: %v\n", info.Size())
	fmt.Printf(" 変更日時 %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf(" ディレクトリ？ %v\n", info.Mode().IsDir())
	fmt.Printf(" 読み書き可能な通常ファイル？ %v\n", info.Mode().IsRegular())
	fmt.Printf(" Unix のファイルアクセス権限ビット %o\n", info.Mode().Perm())
	fmt.Printf(" モードのテキスト表現 %v\n", info.Mode().String())
	if internalStat, ok := info.Sys().(*syscall.Stat_t); ok {
		fmt.Println("Internal Stat")
		fmt.Printf(" デバイス番号 %v\n", internalStat.Dev)
		fmt.Printf(" inode 番号 %v\n", internalStat.Ino)
		fmt.Printf(" モードのテキスト表現 %v\n", internalStat.Mode)
		fmt.Printf(" ハードリンク数 %v\n", internalStat.Nlink)
		fmt.Printf(" 所有者のユーザー ID %v\n", internalStat.Uid)
		fmt.Printf(" 所有者のグループ ID %v\n", internalStat.Gid)
		fmt.Printf(" デバイス ID (特殊ファイルの場合) %v\n", internalStat.Rdev)
		fmt.Printf(" ファイルサイズ %v\n", internalStat.Size)
		fmt.Printf(" ブロックサイズ %v\n", internalStat.Blksize)
		fmt.Printf(" 割り当てられたブロック数 %v\n", internalStat.Blocks)
		//fmt.Printf(" 最終アクセス時刻 %v\n", internalStat.Atimespec)
		//fmt.Printf(" 最終変更時刻 %v\n", internalStat.Mtimespec)
		//fmt.Printf(" 最終状態変更時刻 %v\n", internalStat.Ctimespec)
	}
}
