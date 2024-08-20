// const commitContentBuffer = Buffer.concat(l
// Buffer. from('tree ${this.treeSHA}\n*),
// Buffer. from(parent ${this-parentSHA}\n*),
// Buffer.from
// 'author Piyush Garg <piyushgarg.dev@gmail.com> ${Date.now()} +0000\n*
// Buffer.from
// *committer Piyush Garg <piyushgarg.dev@gmail.com> ${Date.now()}+0000\n\n
// Buffer. from(*${this.message}\n*),
// 1):
// const header = 'commit ${commitContentBuffer.length}\0*;
// const data = Buffer. concat (Buffer. from(header), commitContentBuffer);

package git

import (
	"go-git/pkg/commands/git/internal/blob"
	"go-git/pkg/commands/git/internal/hash"
	"go-git/pkg/utils"
	"os"
	"strings"
	"time"
)

type CommitTreeCommand struct {
	treeSHA   string
	parentSHA string
	message   string
}

func NewCommitTreeCommand() *CommitTreeCommand {
	treeSHA := os.Args[2]
	parentSHA := os.Args[4]
	message := os.Args[6]

	if len(treeSHA) == 0 {
		utils.ErrorLogger("go-git commit-tree: missing tree\n")
		os.Exit(1)
	}

	if len(parentSHA) == 0 {
		utils.ErrorLogger("go-git commit-tree: missing parent\n")
		os.Exit(1)
	}

	if len(message) == 0 {
		utils.ErrorLogger("go-git commit-tree: missing message\n")
		os.Exit(1)
	}

	return &CommitTreeCommand{
		treeSHA:   treeSHA,
		parentSHA: parentSHA,
		message:   message,
	}
}

func (c *CommitTreeCommand) Execute() {

	var commitContentBuffer strings.Builder

	commitContentBuffer.WriteString("tree ")
	commitContentBuffer.WriteString(c.treeSHA)
	commitContentBuffer.WriteString("\n")
	commitContentBuffer.WriteString("parent ")
	commitContentBuffer.WriteString(c.parentSHA)
	commitContentBuffer.WriteString("\n")
	commitContentBuffer.WriteString("author AdarshMishra <adarshmishra969@gmail.com> ")
	commitContentBuffer.WriteString(time.Now().Format(time.RFC3339))
	commitContentBuffer.WriteString(" +0000\n")
	commitContentBuffer.WriteString("committer AdarshMishra <adarshmishra969@gmail.com> ")
	commitContentBuffer.WriteString(time.Now().Format(time.RFC3339))
	commitContentBuffer.WriteString(" +0000\n\n")
	commitContentBuffer.WriteString(c.message)

	var commit strings.Builder
	commit.WriteString("commit ")
	commit.WriteString(commitContentBuffer.String())
	commit.WriteString("\x00")
	commit.WriteString((commitContentBuffer.String()))

	sha1Hash := hash.CreateHash(commit.String())

	blob.WriteBlob(commit.String(), sha1Hash, "")
}
