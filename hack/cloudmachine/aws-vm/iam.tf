data "template_file" "this" {
  template = "${var.iam_policy}"
  vars { }
}

resource "aws_iam_policy" "this" {
  name_prefix = "dev-"
  path        = "/dev/${var.cluster_name}/"
  policy      = "${data.template_file.this.rendered}"
}

resource "aws_iam_role" "this" {
  name_prefix = "dev-"
  path        = "/dev/${var.cluster_name}/"
  assume_role_policy = "${var.iam_role}"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_iam_policy_attachment" "kubernaut-attach" {
  name       = "iam-attach-${var.cluster_name}"
  roles      = ["${aws_iam_role.this.name}"]
  policy_arn = "${aws_iam_policy.this.arn}"
}

resource "aws_iam_instance_profile" "this" {
  name_prefix = "dev-"
  path = "/dev/${var.cluster_name}/"
  role = "${aws_iam_role.this.name}"

  lifecycle {
    create_before_destroy = true
  }
}
