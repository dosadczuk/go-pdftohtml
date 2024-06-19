// Package pdftohtml is a wrapper aroung Xpdf command line tool `pdftohtml`.
//
// What is `pdftohtml`?
//
//	Pdftohtml converts Portable Document Format (PDF) files to HTML.
//
// Reference: https://www.xpdfreader.com/pdftohtml-man.html
package pdftohtml

import (
	"context"
	"os/exec"
	"strconv"
)

// ----------------------------------------------------------------------------
// -- `pdftohtml`
// ----------------------------------------------------------------------------

type command struct {
	path string
	args []string
}

// NewCommand creates new `pdftohtml` command.
func NewCommand(opts ...option) *command {
	cmd := &command{path: "/usr/bin/pdftohtml"}
	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

// Run executes prepared `pdftohtml` command.
func (c *command) Run(ctx context.Context, inpath, outdir string) error {
	cmd := exec.CommandContext(ctx, c.path, append(c.args, inpath, outdir)...)

	return cmd.Run()
}

// String returns a human-readable description of the command.
func (c *command) String() string {
	return exec.Command(c.path, append(c.args, "<inpath>", "<outdir>")...).String()
}

// ----------------------------------------------------------------------------
// -- `pdftohtml` options
// ----------------------------------------------------------------------------

type option func(*command)

// Set custom location for `pdftotext` executable.
func WithCustomPath(path string) option {
	return func(c *command) {
		c.path = path
	}
}

// Read config-file in place of ~/.xpdfrc or the system-wide config file.
func WithCustomConfig(path string) option {
	return func(c *command) {
		c.args = append(c.args, "-cfg", path)
	}
}

// This option tells pdftohtml to instead overwrite the existing directory.
//
// By default pdftohtml will not overwrite the output directory. If the directory already
// exists, pdftohtml will exit with an error.
func WithOutdirOverwrite() option {
	return func(c *command) {
		c.args = append(c.args, "-overwrite")
	}
}

// Specifies the first page to convert.
func WithPageFrom(page uint64) option {
	return func(c *command) {
		c.args = append(c.args, "-f", strconv.FormatUint(page, 10))
	}
}

// Specifies the last page to convert.
func WithPageTo(page uint64) option {
	return func(c *command) {
		c.args = append(c.args, "-l", strconv.FormatUint(page, 10))
	}
}

// Specifies the range of pages to convert.
func WithPageRange(from, to uint64) option {
	return func(c *command) {
		WithPageFrom(from)
		WithPageTo(to)
	}
}

// Specifies the initial zoom level.
//
// The default is 1.0, which means 72dpi, i.e., 1 point in the PDF file will
// be 1 pixel in the HTML.
//
// Using ´-z 1.5’, for example, will make the initial view 50% larger.
func WithInitialZoom(zoom float64) option {
	return func(c *command) {
		c.args = append(c.args, "-z", strconv.FormatFloat(zoom, 'e', 2, 64))
	}
}

// Specifies the resolution, in DPI, for background images. This controls the
// pixel size of the background image files.
//
// The initial zoom level is set by the `WithInitialZoom` option. Specifying
// a larger zoom value will allow the viewer to zoom in farther without upscaling
// artifacts in the background.
func WithResolution(dpi uint64) option {
	return func(c *command) {
		c.args = append(c.args, "-r", strconv.FormatUint(dpi, 10))
	}
}

// Specifies a vertical stretch factor.
//
// Setting this to a value greater than 1.0 will stretch each page vertically,
// spreading out the lines. This also stretches the background image to match.
func WithVerticalStretch(factor float64) option {
	return func(c *command) {
		c.args = append(c.args, "-vstretch", strconv.FormatFloat(factor, 'e', 2, 64))
	}
}

// Embeds the background image as base64-encoded data directly in the HTML file,
// rather than storing it as a separate file.
func WithEmbededBackground() option {
	return func(c *command) {
		c.args = append(c.args, "-embedbackground")
	}
}

// Disable extraction of embedded fonts.
//
// By default, pdftohtml extracts TrueType and OpenType fonts. Disabling extraction
// can work around problems with buggy fonts.
func WithNoFonts() option {
	return func(c *command) {
		c.args = append(c.args, "-nofonts")
	}
}

// Embeds any extracted fonts as base64-encoded data directly in the HTML file, rather
// than storing them as separate files.
func WithEmbededFonts() option {
	return func(c *command) {
		c.args = append(c.args, "-embedfonts")
	}
}

// Don’t draw invisible text.
//
// By default, invisible text (commonly used in OCR’ed PDF files) is drawn as transparent
// (alpha=0) HTML text. This option tells pdftohtml to discard invisible text entirely.
func WithNoInvisibleText() option {
	return func(c *command) {
		c.args = append(c.args, "-skipinvisible")
	}
}

// Treat all text as invisible.
//
// By default, regular (non-invisible) text is not drawn in the background image, and is
// instead drawn with HTML on top of the image. This option tells pdftohtml to include the
// regular text in the background image, and then draw it as transparent (alpha=0) HTML text.
func WithAllInvisibleText() option {
	return func(c *command) {
		c.args = append(c.args, "-allinvisible")
	}
}

// Convert AcroForm text and checkbox fields to HTML input elements.
//
// This also removes text (e.g., underscore characters) and erases background image content
// (e.g., lines or boxes) in the field areas.
func WithFormFields() option {
	return func(c *command) {
		c.args = append(c.args, "-formfields")
	}
}

// Include PDF document metadata as ’meta’ elements in the HTML header.
func WithMeta() option {
	return func(c *command) {
		c.args = append(c.args, "-meta")
	}
}

// Use table mode when performing the underlying text extraction.
//
// This will generally produce better output when the PDF content is a full-page table.
//
// Note: This does not generate HTML tables; it just changes the way text is split up.
func WithModeTable() option {
	return func(c *command) {
		c.args = append(c.args, "-table")
	}
}

// Specify the owner password for the PDF file.
//
// Providing this will bypass all security restrictions.
func WithOwnerPassword(password string) option {
	return func(c *command) {
		c.args = append(c.args, "-opw", password)
	}
}

// Specify the user password for the PDF file.
func WithUserPassword(password string) option {
	return func(c *command) {
		c.args = append(c.args, "-upw", password)
	}
}
