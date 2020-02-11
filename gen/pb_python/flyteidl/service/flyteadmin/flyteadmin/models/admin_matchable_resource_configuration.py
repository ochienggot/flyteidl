# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_matching_attributes import AdminMatchingAttributes  # noqa: F401,E501


class AdminMatchableResourceConfiguration(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'attributes': 'AdminMatchingAttributes',
        'domain': 'str',
        'project': 'str',
        'workflow': 'str'
    }

    attribute_map = {
        'attributes': 'attributes',
        'domain': 'domain',
        'project': 'project',
        'workflow': 'workflow'
    }

    def __init__(self, attributes=None, domain=None, project=None, workflow=None):  # noqa: E501
        """AdminMatchableResourceConfiguration - a model defined in Swagger"""  # noqa: E501

        self._attributes = None
        self._domain = None
        self._project = None
        self._workflow = None
        self.discriminator = None

        if attributes is not None:
            self.attributes = attributes
        if domain is not None:
            self.domain = domain
        if project is not None:
            self.project = project
        if workflow is not None:
            self.workflow = workflow

    @property
    def attributes(self):
        """Gets the attributes of this AdminMatchableResourceConfiguration.  # noqa: E501


        :return: The attributes of this AdminMatchableResourceConfiguration.  # noqa: E501
        :rtype: AdminMatchingAttributes
        """
        return self._attributes

    @attributes.setter
    def attributes(self, attributes):
        """Sets the attributes of this AdminMatchableResourceConfiguration.


        :param attributes: The attributes of this AdminMatchableResourceConfiguration.  # noqa: E501
        :type: AdminMatchingAttributes
        """

        self._attributes = attributes

    @property
    def domain(self):
        """Gets the domain of this AdminMatchableResourceConfiguration.  # noqa: E501


        :return: The domain of this AdminMatchableResourceConfiguration.  # noqa: E501
        :rtype: str
        """
        return self._domain

    @domain.setter
    def domain(self, domain):
        """Sets the domain of this AdminMatchableResourceConfiguration.


        :param domain: The domain of this AdminMatchableResourceConfiguration.  # noqa: E501
        :type: str
        """

        self._domain = domain

    @property
    def project(self):
        """Gets the project of this AdminMatchableResourceConfiguration.  # noqa: E501


        :return: The project of this AdminMatchableResourceConfiguration.  # noqa: E501
        :rtype: str
        """
        return self._project

    @project.setter
    def project(self, project):
        """Sets the project of this AdminMatchableResourceConfiguration.


        :param project: The project of this AdminMatchableResourceConfiguration.  # noqa: E501
        :type: str
        """

        self._project = project

    @property
    def workflow(self):
        """Gets the workflow of this AdminMatchableResourceConfiguration.  # noqa: E501


        :return: The workflow of this AdminMatchableResourceConfiguration.  # noqa: E501
        :rtype: str
        """
        return self._workflow

    @workflow.setter
    def workflow(self, workflow):
        """Sets the workflow of this AdminMatchableResourceConfiguration.


        :param workflow: The workflow of this AdminMatchableResourceConfiguration.  # noqa: E501
        :type: str
        """

        self._workflow = workflow

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminMatchableResourceConfiguration, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminMatchableResourceConfiguration):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other