package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	sparta "github.com/mweagle/Sparta"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
	"github.com/rs/zerolog"
)

/*
Supported signatures

• func ()
• func () error
• func (TIn), error
• func () (TOut, error)
• func (context.Context) error
• func (context.Context, TIn) error
• func (context.Context) (TOut, error)
• func (context.Context, TIn) (TOut, error)
*/

// Standard AWS λ function

func helloWorld(ctx context.Context) (string, error) {
	logger, loggerOk := ctx.Value(sparta.ContextKeyLogger).(*zerolog.Logger)
	if loggerOk {
		logger.Info().Msg("Accessing structured logger 🙌")
	}
	contextLogger, contextLoggerOk := ctx.Value(sparta.ContextKeyRequestLogger).(*zerolog.Logger)
	if contextLoggerOk {
		contextLogger.Info().Msg("Accessing request-scoped log, with request ID field")
	} else if loggerOk {
		logger.Warn().Msg("Failed to access scoped logger")
	} else {
		fmt.Printf("Failed to access any logger")
	}
	return "Welcome to AWS Lambda! 🙌🎉🍾", nil
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	lambdaFn, _ := sparta.NewAWSLambda("Hello World",
		helloWorld,
		sparta.IAMRoleDefinition{})

	sess := session.Must(session.NewSession())
	awsName, awsNameErr := spartaCF.UserAccountScopedStackName("MyOCIStack",
		sess)
	if awsNameErr != nil {
		fmt.Print("Failed to create stack name\n")
		os.Exit(1)
	}

	// Sanitize the name so that it doesn't have any spaces
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	err := sparta.MainEx(awsName,
		"Simple Sparta application that demonstrates core functionality",
		lambdaFunctions,
		nil,
		nil,
		nil,
		false)
	if err != nil {
		os.Exit(1)
	}
}
