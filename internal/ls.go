package internal

import (
	"fmt"
	"os"
	"os/user"
	"syscall"

	"github.com/jedib0t/go-pretty/v6/table"
)

func ListDir(dir string) error {
	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	finfo, err := f.Readdir(0)
	if err != nil {
		return err
	}

	var content []table.Row

	for _, item := range finfo {
		stat, ok := item.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("cannot get stat info for %s", item.Name())
		}
		owner, err := user.LookupId(fmt.Sprintf("%v", stat.Uid))
		if err != nil {
			owner = &user.User{Username: fmt.Sprintf("%d", stat.Uid)}
		}
		ownerGrp, err := user.LookupGroupId(fmt.Sprintf("%v", stat.Gid))
		if err != nil {
			ownerGrp = &user.Group{Name: fmt.Sprintf("%d", stat.Gid)}
		}

		kind := "5_file"
		ft := item.Mode().Type()
		if ft&os.ModeDir != 0 {
			kind = "1_dir"
		} else if ft&os.ModeSymlink != 0 {
			kind = "2_symlink"
		} else if ft&os.ModeType != 0 {
			kind = "3_special"
		} else if item.Mode().Perm()&0111 != 0 {
			kind = "4_exec"
		}

		row := table.Row{
			kind,
			fmt.Sprintf("%04o", item.Mode().Perm()),
			owner.Username,
			ownerGrp.Name,
			ByteSize(uint64(item.Size())),
			item.ModTime().Format("Jan _2 15:04"),
			item.Name(),
		}
		content = append(content, row)

	}

	t := Newtable()
	t.AppendRows(content)
	t.Render()

	return nil
}
