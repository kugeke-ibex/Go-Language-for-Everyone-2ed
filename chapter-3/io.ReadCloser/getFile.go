package stretcher

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"url"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getFile(u *url.URL) (io.ReadCloser, error) {
	return os.Open(u.Path)
}

func getHTTP(u *url.URL) (io.ReadCloser, error) {
	resp, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Stretcher/"+Version)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil

	func getS3(u *url.URL) (io.ReadCloser, error) {
		svc := s3.New(session.Must(session.NewSession()))
		key := strings.TrimPrefix(u.Path, "/")
		result, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(u.Host),
			Key:    aws.String(key),
		})
		if err != nil {
			return nil, err
		}
		return result.Body, nil
	}
}
