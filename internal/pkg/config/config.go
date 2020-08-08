package config

import (
	"context"

	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/logging"
	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/output"
)

type State struct {
	Output <-chan *output.Message
	Ctx    context.Context
}

// Config represents a Snaffler and ShnafflerPoint configuration options
type Config struct {
	// threading (only here for compatibility with snaffler)
	ShareThreads  int `json:"share_threads,omitempty"`
	TreeThreads   int `json:"tree_threads,omitempty"`
	FileThreads   int `json:"file_threads,omitempty"`
	MaxShareQueue int `json:"max_share_queue,omitempty"`
	MaxTreeQueue  int `json:"max_tree_queue,omitempty"`
	MaxFileQueue  int `json:"max_file_queue,omitempty"`

	// logging
	LogToFile      bool   `json:"log_to_file,omitempty"`
	LogFilePath    string `json:"log_file_path,omitempty"`
	LogToConsole   bool   `json:"log_to_console,omitempty"`
	LogLevelString string `json:"log_level_string,omitempty"`

	// targeting (only here for compatibility with snaffler)
	TargetDomain             string   `json:"target_domain,omitempty"`
	TargetDC                 string   `json:"target_dc,omitempty"`
	ShareFinderEnabled       bool     `json:"share_finder_enabled,omitempty"`
	ComputerTargets          []string `json:"computer_targets,omitempty"`
	PathTargets              []string `json:"path_targets,omitempty"`
	ScanSysvol               bool     `json:"scan_sysvol,omitempty"`
	ScanNetlogon             bool     `json:"scan_netlogon,omitempty"`
	QueryDomainForUsers      bool     `json:"query_domain_for_users,omitempty"`
	DomainUsersToMatch       []string `json:"domain_users_to_match,omitempty"`
	DomainUsersWordlistRules []string `json:"domain_users_wordlist_rules,omitempty"`
	MaxSizeToGrep            int      `json:"max_size_to_grep,omitempty"`
	MatchContextBytes        int      `json:"match_context_bytes,omitempty"`

	// snaffling options (our bread and butter)
	Snaffle          bool             `json:"snaffle,omitempty"`
	MaxSizeToSnaffle int              `json:"max_size_to_snaffle,omitempty"`
	SnafflePath      string           `json:"snaffle_path,omitempty"`
	ClassifierRules  []ClassifierRule `json:"classifier_rules,omitempty"`

	// the config rules for this programme
	SharepointURL              string                  `json:"sharepoint_url,omitempty"`
	UseSharepointSiteAllowList bool                    `json:"use_sharepoint_site_allow_list,omitempty"`
	SharepointSiteAllowList    string                  `json:"sharepoint_site_allow_list,omitempty"`
	Credentials                []SharepointCredentials `json:"credentials,omitempty"`
}

// SharepointCredentials is the set of credentials or stolen tokens required to access sharepoint
type SharepointCredentials struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
}

// ClassifierRule represents a single entry in the chainable classifier system
type ClassifierRule struct {
	EnumerationScope string   `json:"enumeration_scope,omitempty"`
	RuleName         string   `json:"rule_name,omitempty"`
	MatchAction      string   `json:"match_action,omitempty"`
	RelayTarget      string   `json:"relay_target,omitempty"`
	Description      string   `json:"description,omitempty"`
	MatchLocation    string   `json:"match_location,omitempty"`
	WordListType     string   `json:"word_list_type,omitempty"`
	WordList         []string `json:"word_list,omitempty"`
	Triage           string   `json:"triage,omitempty"`
}

// Parse a config file contents into a usable object!!!!11!!1
func Parse() *Config {
	// TODO
	logging.Warn("NOT IMPLEMENTED")
	return &Config{
		ShareThreads:               0,
		TreeThreads:                0,
		FileThreads:                0,
		MaxShareQueue:              0,
		MaxTreeQueue:               0,
		MaxFileQueue:               0,
		LogToFile:                  false,
		LogFilePath:                "",
		LogToConsole:               false,
		LogLevelString:             "",
		TargetDomain:               "",
		TargetDC:                   "",
		ShareFinderEnabled:         false,
		ComputerTargets:            nil,
		PathTargets:                nil,
		ScanSysvol:                 false,
		ScanNetlogon:               false,
		QueryDomainForUsers:        false,
		DomainUsersToMatch:         nil,
		DomainUsersWordlistRules:   nil,
		MaxSizeToGrep:              0,
		MatchContextBytes:          0,
		Snaffle:                    false,
		MaxSizeToSnaffle:           0,
		SnafflePath:                "",
		ClassifierRules:            nil,
		SharepointURL:              "",
		UseSharepointSiteAllowList: false,
		SharepointSiteAllowList:    "",
		Credentials:                nil,
	}
}
