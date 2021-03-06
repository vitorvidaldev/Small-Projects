# IAM Section

IAM stands for Identity and Access Management.
IAM is a global service, not region based.

AWS Identity and Access Management (IAM) is a web service for securely controlling access to AWS resources. It enables you to create and control services for user authentication or limit access to a certain set of people who use your AWS resources.

AWS Identity and Access Management (IAM) is a web service that helps you securely control access to AWS resources. You use IAM to control who is authenticated (signed in) and authorized (has permissions) to use resources.

Free to use: AWS Identity and Access Management (IAM) and AWS Security Token Service (AWS STS) are features of your AWS account offered at no additional charge. You are charged only when you access other AWS services using your IAM users or AWS STS temporary security credentials.

## TLDR

IAM (Identity and Access Management) is a global service, resposible for controlling access to AWS resources. It's used to control authentication or to limit access to a set of users.

Users can be grouped, with a feature called Groups. Both users and groups can have policies assigned to them. These policies define what permissions and access they have. Since AWS applies the Least Privilege Principle, a given user will have access based on the permissions assigned to them. These policies can be represented with a JSON document.

## Users & Groups

- Root accounts created by default, they shouldn't be used or shared.
- Users are people within your organization, and they can be grouped.
- Groups only contain users, not other groups.
- Users don't have to belong to a group, and user can belong to multiple groups.

## Permissions

- Users or Groups can be assigned JSON documents called policies.
- These policies define the permissions of the users.
- In AWS, you apply the least privilege principle: don't give more permissions than a user needs.

## Management Console for IAM (Identity and Access Management)

- Access management:
- - User groups: View and configure groups, each having specific permissions.
- - Users: View and configure users.
- - Roles:
- - Policies: List of actions allowed, per user or group
- - Identity providers:

## Policies inheritance

Policies structure:
- Version: policy language version, always include "2012-10-17"
- Id: an identifier for the policy (optional)
- Statement: one or more individual statements (required)

Statements structure:
- Sid: an identifier for the statement (optional)
- Effect: whether the statement allows or denies access (Allow, Deny)
- Principal: account/user/role to which this policy applied to
- Action: list of actions this policy allows or denies
- Condition: conditions for when this policy is in effect (optional)

A statement in an IAM Policy consists of Sid, Effect, Principal, Action, Resource, and Condition. Version is part of the IAM Policy itself, not the statement.

## Password Policy

You can defined a password policy, making sure you have strong passwords and a higher security for your account.

With a password policy, you can:
- Set a minimum password length;
- Require specific character types;
- Allow all IAM users to change their own passwords
- Require users to change their password after some time (password expiration)
- Prevent password re-use

## Multi Factor Authentication (MFA)

Users have access to your account and can possibly change configurations or delete resources in your AWS account. To protect your Root Accounts and IAM users, you can setup multi factor authentication.

MFA = password you know + security device you own

Main benefit of MFA: If a password is stolen or hacked, the account is not compromised.

Security device options:
- Virtual MFA device (Google Authenticator, Authy). Has support for multiple tokens on a single device.
- Universal 2nd Factor (U2F) Security Key. Has support for multiple root and IAM users using a single security key.
- Hardware Key Fov MFA Device
- Hardware Key Fob MFA Device for AWS GovCloud.

## How can users access AWS?

- AWS Management Console (protected by password + MFA)
- AWS Command Line Interface (CLI): protected by access keys
- AWS Software Developer Kit (SDK) - for code: protected by access keys

- Access Keys are generated through the AWS Console
- Users manage their own acces keys
- Access Keys are secret, just like a password. Don't share them
- Access Key ID ~= username
- Secret Access Key ~= password

Do not share your access keys

## AWS CLI

The AWS CLI is a tool that enables you to interact with AWS services using commands in your command-line shell. It gives you direct access to the public APIs of AWS services, and you can develop scripts to manage resources through the CLI.

The AWS CLI is open source, and can be accessed [here](https://github.com/aws/aws-cli/tree/v2).

## AWS SDK

AWS provides Software Development Kits (AWS SDK), that are language-specific APIs (set of libraries) that enable you to access and manage AWS services programmatically.

The AWS SDK is used embedded in your application.

## AWS CloudShell

It gives you a command line access to AWS resources and tools directly from a browser.

AWS CloudShell is a browser-based shell that makes it easy to securely manage, explore, and interact with you AWS resources. CloudShell is pre-authenticated with your console credentials. Common development and operations tools are pre-installed, so no local installation or configuration is required. With CloudShell, you can quickly run scripts with the AWS Command Line Interface (CLI), experiment with AWS service APIs using the AWS SDKs, or use a range of other tools to be productive. You can use CloudShell right from you browser and at no additional cost.

Benefits:
- No extra credentials to manage
- Automatically updated
- No cost
- Customizable

## IAM (Identity and Access Management) Roles for Services

- Some AWS services will need to perform actions on your behalf;
- To do so, we will assign permissions to AWS services with IAM Roles;

Common Roles:
- EC2 Instance Roles
- Lambda Function Roles
- Roles for CloudFormation

An IAM role is an IAM identity that you can create in your account that has specific permissions. An IAM role is similar to an IAM user, in that it is an AWS identity with permission policies that determine what the identity can and cannot do in AWS. However, instead of being uniquely associated with one person, a role is intended to be assumable by anyone who needs it. Also, a role does not have standard long-term credentials such as password or access keys associated with it. Instead, when you assume a role, it provides you with temporary security credentials for your role session.

## IAM Security Tools

IAM Credentials Report (account-level): A report that lists all your account's users and the status of their various credentials.
IAM Access Advisor (user-level): Access advisor shows the service permissions granted to a user and when those services were last accessed.

## IAM Best Practices

- Don't use the root account except for AWS account setup;
- One physical user = One AWS user;
- Assign users to groups and assign permissions to groups;
- Create a strong password policy;
- Use and enforce the use of Multi Factor Authentication (MFA);
- Create and use Roles for giving permissions to AWS services;
- Use Access Keys for Programmatic Access (CLI/SDK);
- Audit permissions of your account with the IAM Credentials Report;
- Never share IAM users & Access Keys.

## IAM Terms

IAM Resources: The user, group, role, policy, and identity provider objects that are stored in IAM. As with other AWS services, you can add, edit, and remove resources from IAM.

IAM Identities: The IAM resource objects that are used to identify and group. You can attach a policy to an IAM identity. Thse include users, groups, and roles.

IAM Entities: The IAM resource objects that AWS uses for authentication. These include IAM users and roles.

Principals: A person or application that uses the AWS account root user, and IAM user, or an IAM role to sign in and make requests to AWS. Principals include federated users and assumed roles.

## IAM Policies

Most policiesa are stored in AWS as JSON documents. You can use the visual editor in the AWS Management Console to create and edit customer managed policies without ever using JSON.

A JSON policy document includes these elements:
- Optional policy-wide information at the top of the document
- One or more individual statements

Each statement includes information about a single permission. If a policy includes multiple statements, AWS applies a logical OR across the statements when evaluating them. If multiple policies apply to a request, AWS applies a logical OR across all of those policies when evaluating them.

The information in a statement is contained within a series of elements.
- Version: Specify the version of the policy language that you want to use.
- Statement: Use this main policy element as a container for the following elements. You can include more than one statement in a policy.
- Sid (Optional): Include an optional statement ID to differentiate between your statements.
- Effect: Use Allow or Deny to indicate whether the policy allows or denies access.
- Principal (Required in only some circumstances) - If you create a resource-based policy, you must indicate the account, user, role, or federated user to which you would like to allow or deny access. If you are creating an IAM permissions policy to attach to a user or role, you cannot include this element. The principal is implied as that user or role.
- Action: Includes a list of actions that the policy allows or denies.
- Resource (Required in only some circunstances): If you create an IAM permissions policy, you must specify a list of resources to which the actions apply. If you create a resource-based policy, this element is optional. If you do not include this element, then the resource to which the action applies is the resource to which the policy is attached.


Example:
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "FirstStatement",
      "Effect": "Allow",
      "Action": ["iam:ChangePassword"],
      "Resource": "*"
    },
    {
      "Sid": "SecondStatement",
      "Effect": "Allow",
      "Action": "s3:ListAllMyBuckets",
      "Resource": "*"
    },
    {
      "Sid": "ThirdStatement",
      "Effect": "Allow",
      "Action": [
        "s3:List*",
        "s3:Get*"
      ],
      "Resource": [
        "arn:aws:s3:::confidential-data",
        "arn:aws:s3:::confidential-data/*"
      ],
      "Condition": {"Bool": {"aws:MultiFactorAuthPresent": "true"}}
    }
  ]
}
```