package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"syscall"
)
func countFiles() int {
	count := 0
	filepath.WalkDir(".", func(path string, file fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !file.IsDir() {

			fmt.Println("--- Path: ", path)
			fmt.Println("Base File Name: ", file.Name())

			// Read file info
			//
			fileinfo, err := file.Info()
			fmt.Println("modification time: ", fileinfo.ModTime())
			fmt.Println("length in bytes: ", fileinfo.Size())

			// Stat file
			//
			stat, _ := fileinfo.Sys().(*syscall.Stat_t)
			fmt.Println("Inode number for file", stat.Ino)
			fmt.Println("File type and permissions", stat.Mode)
			fmt.Println("Number of hard links to file", stat.Nlink)
			fmt.Println("Group ID of file owner", stat.Gid)
			fmt.Println("User ID of file owner", stat.Uid)
			fmt.Println("IDs for device special files", stat.Rdev)
			fmt.Println("Number of 512B blocks allocated", stat.Blocks)
			fmt.Println("Optimal block size for I/O in bytes", stat.Blksize)
			fmt.Println("Total file size in bytes", stat.Size)
			fmt.Println("Time of last file access", stat.Atim.Sec, stat.Atim.Nsec)
			fmt.Println("Time of last file modification", stat.Mtim.Sec, stat.Mtim.Nsec )
			fmt.Println("Time of last status change", stat.Ctim.Sec, stat.Ctim.Nsec )

			count++
			return err
		}

		return nil
	})
	return count
}

func main() {
	count := countFiles()
	fmt.Println(count)
}
