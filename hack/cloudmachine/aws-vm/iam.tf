data "template_file" "this" {
  template = "${var.iam_policy}"
  vars { }
}

resource "aws_iam_policy" "this" {
  name_prefix = "policy-"
  path        = "/dev/${random_pet.this.id}/"
  policy      = "${data.template_file.this.rendered}"
}

resource "aws_iam_role" "this" {
  name_prefix = "role-"
  path        = "/dev/${random_pet.this.id}/"
  assume_role_policy = "${var.iam_role}"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_iam_policy_attachment" "kubernaut-attach" {
  name       = "iam-attach-${random_pet.this.id}"
  roles      = ["${aws_iam_role.this.name}"]
  policy_arn = "${aws_iam_policy.this.arn}"
}

resource "aws_iam_instance_profile" "this" {
  name_prefix = "profile-"
  path = "/dev/${random_pet.this.id}/"
  role = "${aws_iam_role.this.name}"

  lifecycle {
    create_before_destroy = true
  }
}
