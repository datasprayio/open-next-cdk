// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk


type BaseSiteReplaceProps struct {
	Files *string `field:"required" json:"files" yaml:"files"`
	Replace *string `field:"required" json:"replace" yaml:"replace"`
	Search *string `field:"required" json:"search" yaml:"search"`
}

