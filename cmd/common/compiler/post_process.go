package compiler

const PostProcessEntrySectionAddress = uint16(0)
const PostProcessEntrySectionStringSize = uint16(1)

type PostProcessEntry struct {
	Type    uint16
	Section string
	Offset  uint64
	Target  string
}

func NewPostProcessEntry(entryType uint16, section string, offset uint64, target string) *PostProcessEntry {
	return &PostProcessEntry{
		Type:    entryType,
		Section: section,
		Offset:  offset,
		Target:  target,
	}
}
