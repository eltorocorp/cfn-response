# cfn-response
Implementation of the AWS CloudFormation cfn-response objects in Go.

See [AWS's CloudFormation custom resource documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) for information about the basis of this package.

This implementation of cfn-response was written from the ground-up, and isn't a port of any other existing projects.

Similar cfn-response solutions exist for other languages. I haven't used any of them, so, don't consider this a recommendation, but, at the time of this writing, the three most popular are:
 - Python: [jorgebastida/cfn-response](https://github.com/jorgebastida/cfn-response)
 - Java: [SunRun/cfn-response-java](https://github.com/SunRun/cfn-response-java)
 - JavaScript: [LukeMizuhashi/cfn-response](https://github.com/LukeMizuhashi/cfn-response)
