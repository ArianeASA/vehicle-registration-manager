import sys
import json
import os

state_file = sys.argv[1] if len(sys.argv) > 1 else 'default.tfstate'
output_name = sys.argv[2] if len(sys.argv) > 2 else 'default_output_name'

with open(state_file, 'r') as file:
    state = json.load(file)

outputs = state.get('outputs', {})
#
# for key, value in outputs.items():
#     print(f"{key}: {value['value']}")

desired_output = outputs.get(output_name, {}).get('value', '')
# print(f"Output : {desired_output}")

env_file = os.getenv('GITHUB_ENV')
if env_file:
    with open(env_file, "a") as myfile:
        myfile.write(f"{output_name.upper()}={desired_output}\n")
else:
    print(f"Error: GITHUB_ENV not found. Cannot set {output_name.upper()} environment variable.")

