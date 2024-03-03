from setuptools import setup

setup(
   name='common',
   version='1.0',
   description='Common utils and logic for local trending jobs',
   author='Kiran',
   author_email='',
   packages=['common'],  #same as name
   install_requires=['google-api-python-client'], #external packages as dependencies
)