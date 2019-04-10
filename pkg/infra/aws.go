package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"net"
	"time"
)

type AWS struct {
	ec2    *ec2.EC2
	Region string
}

type awsVPC struct {
	aws *AWS
}

func (v *awsVPC) ID() string {
	return ""
}

func (v *awsVPC) Name() string {
	return ""
}

type EC2Machine struct {
	// ID is how the machine is referenced. The format and meaning of a Machine ID is opaque to users.
	id string

	// Class indicates what "type" or "size" of machine is allocated.
	class string

	// Flavor identifies the underlying infrastructure provider.
	flavor string

	// Family identifies the underlying product class on the infrastructure provider.
	family string

	// CreationTime is when the instance was started by the infrastructure provider.
	creationTime time.Time

	// Reference to the underlying cloud provider.
	aws *AWS
}

func (m *EC2Machine) ID() string {
	return m.id
}

func (m *EC2Machine) Class() string {
	return m.class
}

func (m *EC2Machine) Flavor() string {
	return m.flavor
}

func (m *EC2Machine) Family() string {
	return m.family
}

func (m *EC2Machine) CreationTime() time.Time {
	return m.creationTime
}

func (m *EC2Machine) PublicIP() net.IP {
	return nil
}

func (m *EC2Machine) PublicDNS() string {
	return ""
}

func (m *EC2Machine) PrivateIP() net.IP {
	return nil
}

func (m *EC2Machine) PrivateDNS() string {
	return ""
}

func (m *EC2Machine) Delete() (bool, error) {
	return m.aws.deleteMachine(m)
}

func (a *AWS) CreateMachine() (*Machine, error) {
	return nil, nil
}

func (a *AWS) GetMachineByID(ID string) (Machine, bool, error) {
	var machine *EC2Machine
	var found bool
	var err error

	q := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{&ID},
	}

	res, err := a.ec2.DescribeInstances(q)
	if err != nil {
		return nil, found, err
	}

	var instance *ec2.Instance
	for _, reservation := range res.Reservations {
		if len(reservation.Instances) >= 1 {
			instance = reservation.Instances[0]
			break
		}
	}

	if instance == nil {
		return nil, found, nil
	}

	found = true
	machine = &EC2Machine{
		id:           aws.StringValue(instance.InstanceId),
		class:        aws.StringValue(instance.InstanceType),
		flavor:       "aws",
		family:       "ec2",
		creationTime: aws.TimeValue(instance.LaunchTime),
		aws:          nil,
	}

	return machine, found, nil
}

func (a *AWS) deleteMachine(m *EC2Machine) (bool, error) {
	q := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{&m.id},
	}

	// TODO: This is probably wrong
	_, err := a.ec2.TerminateInstances(q)
	if err != nil {
		return false, err
	}

	return true, nil
}
