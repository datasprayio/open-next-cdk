// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk


type CreateArchiveArgs struct {
	Directory *string `field:"required" json:"directory" yaml:"directory"`
	ZipFileName *string `field:"required" json:"zipFileName" yaml:"zipFileName"`
	ZipOutDir *string `field:"required" json:"zipOutDir" yaml:"zipOutDir"`
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	FileGlob *string `field:"optional" json:"fileGlob" yaml:"fileGlob"`
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
}

