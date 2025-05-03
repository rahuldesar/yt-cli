package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMapLater struct {
	Up             key.Binding
	Down           key.Binding
	FirstItem      key.Binding
	LastItem       key.Binding
	TogglePreview  key.Binding
	Refresh        key.Binding
	PageDown       key.Binding
	PageUp         key.Binding
	NextSection    key.Binding
	PrevSection    key.Binding
	DownloadVideo  key.Binding
	DownloadAudio  key.Binding
	Help           key.Binding
	Quit           key.Binding
	ChannelSearch  key.Binding
	PlaylistSearch key.Binding
	HistorySearch  key.Binding
	RecentSearch   key.Binding
	// TODO: MAYBE THESE ARE GOOD
	OpenYoutube key.Binding
}

type Dimensions struct {
	Width  int
	Height int
}

const (
	WaitingIcon = ""
	FailureIcon = "󰅙"
	SuccessIcon = ""

	CommentIcon = ""
	DraftIcon   = ""
	BehindIcon  = "󰇮"
	BlockedIcon = ""
	MergedIcon  = ""
	OpenIcon    = ""
	ClosedIcon  = ""

	NewContributorIcon = "󰎔" // \udb80\udf94 nf-md-new_box
	ContributorIcon    = "" // \uedc6 nf-fa-user_check
	CollaboratorIcon   = "" // \uedcf nf-fa-user_shield
	MemberIcon         = "" // \uf42b nf-oct-organization
	OwnerIcon          = "󱇐" // \udb84\uddd0 nf-md-crown_outline
	UnknownRoleIcon    = "󱐡" // \udb85\udc21 nf-md-incognito_circle

	// ME : NEWLY ADDED ICONS
	SearchIcon1 = " "
	SearchIcon2 = " "
	SearchIcon3 = " "

	YoutubeIcon = " "

	CalendarIcon = " " //      
	PersonIcon   = " " //   - Used for Uploader
	ClockIcon    = " " // 󰥔
	ThumbsUpIcon = " " // 
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),        // actual keybindings
		key.WithHelp("↑/k", "move up"), // corresponding help text
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
}
