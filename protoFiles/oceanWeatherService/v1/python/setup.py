from setuptools import setup

setup(
    # Needed to silence warnings (and to be a worthwhile package)
    name='OceanWeatherServiceProtos',
    url='https://github.com/NicholasBunn/mastersCaseStudy/protoFiles/oceanWeatherService/v1/python',
    author='Nichola Bunn',
    author_email='nicholasbunn04@gmail.com',
    # Needed to actually package something
    packages=['measure'],
    # Needed for dependencies
    install_requires=['google.protobuf', 'grpc'],
    # *strongly* suggested for sharing
    version='1',
    # The license can be anything you like
    license='MIT',
    description='Python\'s generated proto files for the Ocean Weather Service',
    # We will also need a readme eventually (there will be a warning)
    # long_description=open('README.txt').read(),
)