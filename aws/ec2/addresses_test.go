package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Addresses", func() {
	var (
		client *fakes.AddressesClient
		logger *fakes.Logger

		addresses ec2.Addresses
	)

	BeforeEach(func() {
		client = &fakes.AddressesClient{}
		logger = &fakes.Logger{}

		addresses = ec2.NewAddresses(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeAddressesCall.Returns.Output = &awsec2.DescribeAddressesOutput{
				Addresses: []*awsec2.Address{{
					PublicIp:     aws.String("banana"),
					AllocationId: aws.String("the-allocation-id"),
					InstanceId:   aws.String(""),
				}},
			}
			filter = "ban"
		})

		It("releases ec2 addresses", func() {
			items, err := addresses.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeAddressesCall.CallCount).To(Equal(1))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to release address banana?"))

			Expect(items).To(HaveLen(1))
			Expect(items).To(HaveKeyWithValue("banana", "the-allocation-id"))
		})

		Context("when the address name does not contain the filter", func() {
			PIt("does not try to release it", func() {
				// The address resource may not be named after the environment
			})
		})

		Context("when the address is in use by an instance", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = true
				client.DescribeAddressesCall.Returns.Output = &awsec2.DescribeAddressesOutput{
					Addresses: []*awsec2.Address{{
						PublicIp:   aws.String("banana"),
						InstanceId: aws.String("the-instance-using-it"),
					}},
				}
			})

			It("does not try to release it", func() {
				items, err := addresses.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeAddressesCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the client fails to describe addresses", func() {
			BeforeEach(func() {
				client.DescribeAddressesCall.Returns.Error = errors.New("some error")
			})

			It("does not try releasing them", func() {
				_, err := addresses.List(filter)
				Expect(err).To(MatchError("Describing addresses: some error"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not release the address", func() {
				items, err := addresses.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to release address banana?"))
				Expect(items).To(HaveLen(0))
			})
		})
	})

	Describe("Delete", func() {
		var items map[string]string

		BeforeEach(func() {
			items = map[string]string{"banana": "the-allocation-id"}
		})

		It("releases ec2 addresses", func() {
			err := addresses.Delete(items)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ReleaseAddressCall.CallCount).To(Equal(1))
			Expect(client.ReleaseAddressCall.Receives.Input.AllocationId).To(Equal(aws.String("the-allocation-id")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS releasing address banana\n"}))
		})

		Context("when the client fails to release the address", func() {
			BeforeEach(func() {
				client.ReleaseAddressCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := addresses.Delete(items)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR releasing address banana: some error\n"}))
			})
		})
	})
})
