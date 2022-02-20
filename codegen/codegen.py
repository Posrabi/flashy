import glob
import os
from matplotlib.pyplot import close
from pandas import array
import yaml

# returns abs path of the parent


def absParent(path):
    return os.path.abspath(os.path.join(path, os.pardir))
    # os.pardir is ..


curDir = absParent(__file__)
rootDir = absParent(curDir)

defaultOutputDir = os.path.join(rootDir, 'backend/api/pkg/api')
defaultInputDir = os.path.join(curDir, 'inputs/api.txt')

ServiceFile = os.path.join(curDir, 'services.yaml')


def parseServices(serviceFile):
    with open(serviceFile) as file:
        services = yaml.load(file, loader=yaml.FullLoader)

        for svc in services.values():
            generate(printOuput="to_file", outputDir=svc["output_dir"],
                     inputFile=svc["input"], inputDirs=svc["input_dirs"], customMappings=svc["custom_mappings"])


def replaceMappings(line: str, customMappings: dict) -> str:
    for k, v in customMappings.items():
        line = line.replace(f'%{k}', v)
    return line


def generate(printOuput="store_true", outputDir=defaultOutputDir, inputFile=defaultInputDir,
             inputDirs=[""], customMappings={}):
    if len(inputDirs) == 0:
        inputDirs.append("")

    endpoints = []
    for line in open(os.path.join(curDir, inputFile)):
        if line:
            endpoints.append(line.strip())
        for inputDir in inputDirs:
            inputDir = os.path.join(rootDir, inputDir)
            for inputFile in glob.glob(os.path.join(inputDir, '*.go.template')):
                inputBase = os.path.basename(inputFile)
                outputFile = os.path.join(
                    rootDir, outputDir, inputBase.replace('.template', ''))

                with open(inputDir) as i:
                    outputLines = []
                    replaceLines = []
                    replace = False

                    for line in i:
                        line = line.rstrip()
                        if line.startswith('[['):
                            replace = True
                        elif line.startswith(']]'):
                            for endpoint in endpoints:
                                endpointCamel = endpoint[0].lower(
                                ) + endpoint[1:]
                                for replaceLine in replaceLines:
                                    outputLines.append(replaceMappings(replaceLine.replace('%s', endpoint)
                                                                       .replace('%l', endpointCamel), customMappings))
                            replace = False
                            replaceLines = []
                        elif replace:
                            replaceLines.append(line)
                        else:
                            outputLines.append(
                                replaceMappings(line, customMappings))
                    if printOuput == 'store_true':
                        print('\n'.join(outputLines))
                    else:
                        with open(outputFile, 'w') as o:
                            o.write('// Code generated .* DO NOT EDIT.\n// To make changes, please modify %s\n\n//nolint\n' %
                                    os.path.join('codegen', inputBase))
                            o.write('\n'.join(outputLines))


if __name__ == '__main__':
    parseServices(ServiceFile)
