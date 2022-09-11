package presentation

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation/icons"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/theme"
)

func GetWorktreeDisplayString(isCurrent bool, isPathMissing bool, worktree *models.Worktree) []string {
	textStyle := theme.DefaultTextColor

	current := ""
	currentColor := style.FgCyan
	if isCurrent {
		current = "  *"
		currentColor = style.FgGreen
	}

	icon := icons.IconForWorktree(worktree, false)
	if isPathMissing {
		textStyle = style.FgRed
		icon = icons.IconForWorktree(worktree, true)
	}

	res := make([]string, 0, 3)
	res = append(res, currentColor.Sprint(current))
	if icons.IsIconEnabled() {
		res = append(res, textStyle.Sprint(icon))
	}
	res = append(res, textStyle.Sprint(worktree.Name()))
	return res
}
