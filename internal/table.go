package internal

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var styleHiddenBorders = table.Style{
	Name: "hidden-borders",
	Box: table.BoxStyle{
		BottomLeft:       "",
		BottomRight:      " ",
		BottomSeparator:  " ",
		Left:             "",
		LeftSeparator:    " ",
		MiddleHorizontal: " ",
		MiddleSeparator:  " ",
		MiddleVertical:   " ",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		Right:            " ",
		RightSeparator:   " ",
		TopLeft:          "",
		TopRight:         " ",
		TopSeparator:     " ",
		UnfinishedRow:    " ",
	},
	Options: table.Options{
		SeparateHeader: true,
	},
}

var header = table.Row{"kind", "perm", "user", "group", "size", "modified", "name"}

var sortByDirectoriesThenByName = []table.SortBy{
	{Name: "kind", Mode: table.AscAlphaNumeric},
	{Name: "name", Mode: table.AscAlphaNumeric},
}

var HideKindColumn = []table.ColumnConfig{
	{
		Name:   "kind",
		Hidden: true,
	},
}

func rowColor(row table.Row) text.Colors {
	kind := row[0]
	switch {
	case kind == "1_dir":
		return text.Colors{text.FgBlue}
	case kind == "2_symlink":
		return text.Colors{text.FgMagenta}
	case kind == "3_special":
		return text.Colors{text.FgRed}
	case kind == "4_exec":
		return text.Colors{text.FgGreen}
	}
	return nil
}

func Newtable() table.Writer {
	t := table.NewWriter()
	t.SetStyle(styleHiddenBorders)
	t.AppendHeader(header)
	t.SortBy(sortByDirectoriesThenByName)
	t.SetRowPainter(rowColor)
	t.SetColumnConfigs(HideKindColumn)
	t.SetOutputMirror(os.Stdout)
	return t
}
