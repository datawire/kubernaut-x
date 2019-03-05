output machine_id {
  value = "${aws_instance.this.id}"
}

output machine_public_ip {
  value = "${aws_instance.this.public_ip}"
}

output machine_public_dns {
  value = "${aws_instance.this.public_dns}"
}

output machine_private_ip {
  value = "${aws_instance.this.private_ip}"
}

output machine_private_dns {
  value = "${aws_instance.this.private_dns}"
}
