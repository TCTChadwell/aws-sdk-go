// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package devicefarm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDeviceFarm_CreateDevicePool() {
	svc := devicefarm.New(nil)

	params := &devicefarm.CreateDevicePoolInput{
		Name:       aws.String("Name"),               // Required
		ProjectARN: aws.String("AmazonResourceName"), // Required
		Rules: []*devicefarm.Rule{ // Required
			{ // Required
				Attribute: aws.String("DeviceAttribute"),
				Operator:  aws.String("RuleOperator"),
				Value:     aws.String("String"),
			},
			// More values...
		},
		Description: aws.String("Message"),
	}
	resp, err := svc.CreateDevicePool(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_CreateProject() {
	svc := devicefarm.New(nil)

	params := &devicefarm.CreateProjectInput{
		Name: aws.String("Name"), // Required
	}
	resp, err := svc.CreateProject(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_CreateUpload() {
	svc := devicefarm.New(nil)

	params := &devicefarm.CreateUploadInput{
		Name:        aws.String("Name"),               // Required
		ProjectARN:  aws.String("AmazonResourceName"), // Required
		Type:        aws.String("UploadType"),         // Required
		ContentType: aws.String("ContentType"),
	}
	resp, err := svc.CreateUpload(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetAccountSettings() {
	svc := devicefarm.New(nil)

	var params *devicefarm.GetAccountSettingsInput
	resp, err := svc.GetAccountSettings(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetDevice() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetDeviceInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevice(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetDevicePool() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetDevicePoolInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetDevicePool(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetDevicePoolCompatibility() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetDevicePoolCompatibilityInput{
		AppARN:        aws.String("AmazonResourceName"), // Required
		DevicePoolARN: aws.String("AmazonResourceName"), // Required
		TestType:      aws.String("TestType"),
	}
	resp, err := svc.GetDevicePoolCompatibility(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetJob() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetJobInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetJob(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetProject() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetProjectInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetProject(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetRun() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetRunInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetRun(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetSuite() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetSuiteInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetSuite(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetTest() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetTestInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetTest(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_GetUpload() {
	svc := devicefarm.New(nil)

	params := &devicefarm.GetUploadInput{
		ARN: aws.String("AmazonResourceName"), // Required
	}
	resp, err := svc.GetUpload(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListArtifacts() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListArtifactsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		Type:      aws.String("ArtifactCategory"),   // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListArtifacts(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListDevicePools() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListDevicePoolsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
		Type:      aws.String("DevicePoolType"),
	}
	resp, err := svc.ListDevicePools(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListDevices() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListDevicesInput{
		ARN:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListDevices(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListJobs() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListJobsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListProjects() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListProjectsInput{
		ARN:       aws.String("AmazonResourceName"),
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListProjects(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListRuns() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListRunsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListRuns(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListSamples() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListSamplesInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSamples(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListSuites() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListSuitesInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListSuites(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListTests() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListTestsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListTests(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListUniqueProblems() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListUniqueProblemsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUniqueProblems(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ListUploads() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ListUploadsInput{
		ARN:       aws.String("AmazonResourceName"), // Required
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListUploads(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleDeviceFarm_ScheduleRun() {
	svc := devicefarm.New(nil)

	params := &devicefarm.ScheduleRunInput{
		AppARN:        aws.String("AmazonResourceName"), // Required
		DevicePoolARN: aws.String("AmazonResourceName"), // Required
		ProjectARN:    aws.String("AmazonResourceName"), // Required
		Test: &devicefarm.ScheduleRunTest{ // Required
			Type:   aws.String("TestType"), // Required
			Filter: aws.String("Filter"),
			Parameters: map[string]*string{
				"Key": aws.String("String"), // Required
				// More values...
			},
			TestPackageARN: aws.String("AmazonResourceName"),
		},
		Configuration: &devicefarm.ScheduleRunConfiguration{
			AuxiliaryApps: []*string{
				aws.String("AmazonResourceName"), // Required
				// More values...
			},
			BillingMethod:       aws.String("BillingMethod"),
			ExtraDataPackageARN: aws.String("AmazonResourceName"),
			Locale:              aws.String("String"),
			Location: &devicefarm.Location{
				Latitude:  aws.Float64(1.0), // Required
				Longitude: aws.Float64(1.0), // Required
			},
			NetworkProfileARN: aws.String("AmazonResourceName"),
			Radios: &devicefarm.Radios{
				Bluetooth: aws.Bool(true),
				Gps:       aws.Bool(true),
				Nfc:       aws.Bool(true),
				Wifi:      aws.Bool(true),
			},
		},
		Name: aws.String("Name"),
	}
	resp, err := svc.ScheduleRun(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}
