package radiantpro

const RadiantToml = `
	dsn = "root:123456@tcp(127.0.0.1:3306)/radiant"
	driver = "mysql"
	proType = "default"
	enableModule = []
	apiPrefix = "/"
	gitRemotePath = "https://github.com/radiant-dev/radiant-pro.git"
	format = true
	sourceGen = "text"
	gitPull = true
	[models.user]
		name = ["uid"]
		orm = ["auto"]
		comment = ["Uid"]
		
`
