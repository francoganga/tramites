package services

import (
	"database/sql"
	"fmt"

	// Required by ent
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/francoganga/go_reno2/config"

	// Require by ent
	_ "github.com/francoganga/go_reno2/ent/runtime"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/uptrace/bun/extra/bundebug"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests
type Container struct {
	// Validator stores a validator
	Validator *Validator

	// Web stores the web framework
	Web *echo.Echo

	// Config stores the application configuration
	Config *config.Config

	// Cache contains the cache client
	Cache *CacheClient

	// Database stores the connection to the database
	Database *sql.DB

	// Bun DB connection
	Bun *bun.DB

	// ORM stores a client to the ORM
	//ORM *ent.Client

	// Mail stores an email sending client
	Mail *MailClient

	// Auth stores an authentication client
	//Auth *AuthClient

	// TemplateRenderer stores a service to easily render and cache templates
	TemplateRenderer *TemplateRenderer

	// Tasks stores the task client
	Tasks *TaskClient
}

// NewContainer creates and initializes a new Container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initValidator()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	//c.initORM()
	//c.initAuth()
	c.initTemplateRenderer()
	c.initMail()
	c.initTasks()
	return c
}

// Shutdown shuts the Container down and disconnects all connections
func (c *Container) Shutdown() error {
	if err := c.Tasks.Close(); err != nil {
		return err
	}
	if err := c.Cache.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}

	return nil
}

// initConfig initializes configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

// initValidator initializes the validator
func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

// initWeb initializes the web framework
func (c *Container) initWeb() {
	c.Web = echo.New()

	// Configure logging
	switch c.Config.App.Environment {
	case config.EnvProduction:
		c.Web.Logger.SetLevel(log.WARN)
	default:
		c.Web.Logger.SetLevel(log.DEBUG)
	}

	c.Web.Validator = c.Validator
}

// initCache initializes the cache
func (c *Container) initCache() {
	var err error
	if c.Cache, err = NewCacheClient(c.Config); err != nil {
		panic(err)
	}
}

// initDatabase initializes the database
// If the environment is set to test, the test database will be used and will be dropped, recreated and migrated
// func (c *Container) initDatabase() {
// 	var err error
//
// 	getAddr := func(dbName string) string {
// 		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
// 			c.Config.Database.User,
// 			c.Config.Database.Password,
// 			c.Config.Database.Hostname,
// 			c.Config.Database.Port,
// 			dbName,
// 		)
// 	}
//
// 	c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.Database))
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to connect to database: %v", err))
// 	}
//
// 	// Check if this is a test environment
// 	if c.Config.App.Environment == config.EnvTest {
// 		// Drop the test database, ignoring errors in case it doesn't yet exist
// 		_, _ = c.Database.Exec("DROP DATABASE " + c.Config.Database.TestDatabase)
//
// 		// Create the test database
// 		if _, err = c.Database.Exec("CREATE DATABASE " + c.Config.Database.TestDatabase); err != nil {
// 			panic(fmt.Sprintf("failed to create test database: %v", err))
// 		}
//
// 		// Connect to the test database
// 		if err = c.Database.Close(); err != nil {
// 			panic(fmt.Sprintf("failed to close database connection: %v", err))
// 		}
// 		c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.TestDatabase))
// 		if err != nil {
// 			panic(fmt.Sprintf("failed to connect to database: %v", err))
// 		}
// 	}
// }

func (c *Container) initDatabase() {
	var err error

	getAddr := func(dbName string) string {
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
			c.Config.Database.User,
			c.Config.Database.Password,
			c.Config.Database.Hostname,
			c.Config.Database.Port,
			dbName,
		)
	}

	c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.Database))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	c.Bun = bun.NewDB(c.Database, pgdialect.New())

	c.Bun.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.FromEnv("BUNDEBUG")))

	// Check if this is a test environment
	if c.Config.App.Environment == config.EnvTest {
		// Drop the test database, ignoring errors in case it doesn't yet exist
		_, _ = c.Database.Exec("DROP DATABASE " + c.Config.Database.TestDatabase)

		// Create the test database
		if _, err = c.Database.Exec("CREATE DATABASE " + c.Config.Database.TestDatabase); err != nil {
			panic(fmt.Sprintf("failed to create test database: %v", err))
		}

		// Connect to the test database
		if err = c.Database.Close(); err != nil {
			panic(fmt.Sprintf("failed to close database connection: %v", err))
		}
		c.Database, err = sql.Open("pgx", getAddr(c.Config.Database.TestDatabase))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}

		c.Bun = bun.NewDB(c.Database, pgdialect.New())
	}
}

// initORM initializes the ORM

// initAuth initializes the authentication client
// func (c *Container) initAuth() {
// 	c.Auth = NewAuthClient(c.Config, c.ORM)
// }

// initTemplateRenderer initializes the template renderer
func (c *Container) initTemplateRenderer() {
	c.TemplateRenderer = NewTemplateRenderer(c.Config)
}

// initMail initialize the mail client
func (c *Container) initMail() {
	var err error
	c.Mail, err = NewMailClient(c.Config, c.TemplateRenderer)
	if err != nil {
		panic(fmt.Sprintf("failed to create mail client: %v", err))
	}
}

// initTasks initializes the task client
func (c *Container) initTasks() {
	c.Tasks = NewTaskClient(c.Config)
}
