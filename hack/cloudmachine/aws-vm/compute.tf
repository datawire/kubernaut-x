resource "tls_private_key" "this" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "this" {
  key_name_prefix = "kubernaut-dev-"
  public_key      = "${tls_private_key.this.public_key_openssh}"
}

resource "aws_instance" "this" {
  ami                         = "${var.image_id}"
  associate_public_ip_address = true
  iam_instance_profile        = "${aws_iam_instance_profile.this.name}"
  instance_type               = "${var.machine_type}"
  subnet_id                   = "${aws_subnet.kubernaut.id}"
  #user_data                   = "${data.template_cloudinit_config.kubernaut_cloudinit.rendered}"
  key_name                    = "${aws_key_pair.this.id}"
  monitoring                  = false
  vpc_security_group_ids      = ["${aws_security_group.kubernaut.id}"]

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "${var.disk_size}"
    delete_on_termination = true
  }

  #tags = "${merge(map("Name", var.cluster_name, "kubernaut.io/cluster/name", var.cluster_name), var.tags)}"
}