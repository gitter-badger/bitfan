// Code generated by "bitfanDoc "; DO NOT EDIT
package tail

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "tail",
  ImportPath: "github.com/vjeantet/bitfan/processors/input-tail",
  Doc:        "",
  DocShort:   "",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "AddField",
        Alias:          "add_field",
        Doc:            "Add a field to an event. Default value is {}",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "CloseOlder",
        Alias:          "close_older",
        Doc:            "Closes any files that were last read the specified timespan in seconds ago.\nDefault value is 3600 (i.e. 1 hour)\nThis has different implications depending on if a file is being tailed or read.\nIf tailing, and there is a large time gap in incoming data the file can be\nclosed (allowing other files to be opened) but will be queued for reopening\nwhen new data is detected. If reading, the file will be closed after\nclose_older seconds from when the last bytes were read.",
        Required:       false,
        Type:           "int",
        DefaultValue:   "3600",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Codec",
        Alias:          "codec",
        Doc:            "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline",
        Required:       false,
        Type:           "codec",
        DefaultValue:   "\"line\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Delimiter",
        Alias:          "delimiter",
        Doc:            "Set the new line delimiter. Default value is \"\\n\"",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"\\n\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "DiscoverInterval",
        Alias:          "discover_interval",
        Doc:            "How often (in seconds) we expand the filename patterns in the path option\nto discover new files to watch. Default value is 15",
        Required:       false,
        Type:           "int",
        DefaultValue:   "15",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Exclude",
        Alias:          "exclude",
        Doc:            "Exclusions (matched against the filename, not full path).\nFilename patterns are valid here, too.",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "IgnoreOlder",
        Alias:          "ignore_older",
        Doc:            "When the file input discovers a file that was last modified before the\nspecified timespan in seconds, the file is ignored.\nAfter it’s discovery, if an ignored file is modified it is no longer ignored\nand any new data is read.\nDefault value is 86400 (i.e. 24 hours)",
        Required:       false,
        Type:           "int",
        DefaultValue:   "86400",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "MaxOpenFiles",
        Alias:          "max_open_files",
        Doc:            "What is the maximum number of file_handles that this input consumes at any one time.\nUse close_older to close some files if you need to process more files than this number.",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Path",
        Alias:          "path",
        Doc:            "The path(s) to the file(s) to use as an input.\nYou can use filename patterns here, such as /var/log/*.log.\nIf you use a pattern like /var/log/**/*.log, a recursive search of /var/log\nwill be done for all *.log files.\nPaths must be absolute and cannot be relative.\nYou may also configure multiple paths.",
        Required:       true,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "SincedbPath",
        Alias:          "sincedb_path",
        Doc:            "Path of the sincedb database file\nThe sincedb database keeps track of the current position of monitored\nlog files that will be written to disk.",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\".sincedb.json\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "SincedbWriteInterval",
        Alias:          "sincedb_write_interval",
        Doc:            "How often (in seconds) to write a since database with the current position of monitored log files.\nDefault value is 15",
        Required:       false,
        Type:           "int",
        DefaultValue:   "15",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "StartPosition",
        Alias:          "start_position",
        Doc:            "Choose where BitFan starts initially reading files: at the beginning or at the end.\nThe default behavior treats files like live streams and thus starts at the end.\nIf you have old data you want to import, set this to beginning.\nThis option only modifies \"first contact\" situations where a file is new\nand not seen before, i.e. files that don’t have a current position recorded in a sincedb file.\nIf a file has already been seen before, this option has no effect and the\nposition recorded in the sincedb file will be used.\nDefault value is \"end\"\nValue can be any of: \"beginning\", \"end\"",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"end\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "StatInterval",
        Alias:          "stat_interval",
        Doc:            "How often (in seconds) we stat files to see if they have been modified.\nIncreasing this interval will decrease the number of system calls we make,\nbut increase the time to detect new log lines.\nDefault value is 1",
        Required:       false,
        Type:           "int",
        DefaultValue:   "1",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Tags",
        Alias:          "tags",
        Doc:            "Add any number of arbitrary tags to your event. There is no default value for this setting.\nThis can help with processing later. Tags can be dynamic and include parts of the event using the %{field} syntax.",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Type",
        Alias:          "type",
        Doc:            "Add a type field to all events handled by this input.\nTypes are used mainly for filter activation.",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}