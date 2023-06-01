---
marp: true
---

# 9. ファイルシステムの基礎と Go 言語の標準パッケージ

kotetsu@nicecode.app

---

## 9.1 ファイルシステムの基礎

- [ファイルシステム比較(Wikipedia)](https://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0#%E6%AF%94%E8%BC%83)
- 読み書きを行う単位=セクター
  - Windows のセクタ情報
    ```dos
    C:\Windows\System32>fsutil fsinfo sectorInfo C:
    LogicalBytesPerSector :                                 512
    PhysicalBytesPerSectorForAtomicity :                    4096
    PhysicalBytesPerSectorForPerformance :                  4096
    FileSystemEffectivePhysicalBytesPerSectorForAtomicity : 4096
    デバイスの配置 :                                        配置 (0x000)
    デバイス上のパーティションの配置 :                      配置 (0x000)
    シーク ペナルティなし
    トリムをサポート
    DAX に対応していません
    ```
- [NAND フラッシュメモリベース SSD のセクタサイズの過去と現在](https://qiita.com/ken-yossy/items/28c2086b1cf02ea3442d)

---

- inode (Linux)

  ```
  ❯ ls -i
  625135 ./   49456 ../  625136 .git/  625165 2_ioWriter/  625182 3_ioReader/   49033 9_basis_of_filesystem/  625228 go.mod
  ```

- vfs (Virtual File System)
  - 様々なファイルシステムを統一的に扱える API
- LVM (Logical Volume Manager)
  - 複数の物理ボリュームをまとめて論理ボリュームとして扱う
  - ボリュームグループを作成し、その中に論理ボリュームを作成する
  - 論理ボリュームはパーティションと同様に扱える

---

9.2 ファイル／ディレクトリを扱う Go 言語の関数たち

- 2,3 章の復習
- sync

  ```
  # SSD の PC
  ❯ go run 9_2_1_sync.go
  Write: 7.505µs
  Sync: 4.884071ms
  Close: 4.466µs

  # HDD(Software raid1) の PC
  ❯ ./9_2_1_sync
  Write: 17.32µs
  Sync: 37.870697ms
  Close: 25.706µs
  ```

  - ディスク遅い
    - https://gist.github.com/hellerbarde/2843375

---

- ファイルモード
  - ファイルモードの表現
    - 8 進数、rwxrwxrwx、rwxr-xr-x、0644、0755
  - ファイルモードの変更
    - chmod コマンド、os.Chmod 関数
  - Go 言語では、ファイルやディレクトリの作成時にこれらの数値を設定
- os.Mkdir / os.MkdirAll
- os.Remove / os.RemoveAll
- os.Truncate
- os.Rename
  - Windows の場合、同一ドライブ内の移動のみ

---

- ファイルの属性と取得
  - os.Stat
  - os.Lstat
- FileMode タイプ
  - https://pkg.go.dev/gopkg.in/src-d/go-git.v4/plumbing/filemode#FileMode
- ファイルの存在チェック
- OS固有のファイル属性
  - os.FileInfoSys()
- ファイルの同一性チェック
  - os.SameFile

---

- ファイルの属性の設定
  - os.Chmod
  - os.Chown
  - os.Chtimes
- リンク
  - os.Link
  - os.Symlink
  - os.Readlink
- ディレクトリ情報の取得
  - os.ReadDir
  - 一覧の取得はがんばる
---
9.3 OS 内部におけるファイル操作の高速化

---

9.4 ファイルパスとマルチプラットフォーム

---

9.5 path/filepath パッケージの関数たち

---

9.6 本章のまとめと次章予告
