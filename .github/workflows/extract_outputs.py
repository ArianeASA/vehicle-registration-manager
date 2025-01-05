import sys
import json
import os

# Obter o nome do arquivo dos argumentos de linha de comando
state_file = sys.argv[1] if len(sys.argv) > 1 else 'default.tfstate'
output_name = sys.argv[2] if len(sys.argv) > 2 else 'default_output_name'

# Carregar o arquivo de estado do Terraform
with open(state_file, 'r') as file:
    state = json.load(file)

# Extrair os outputs
outputs = state.get('outputs', {})

# Imprimir os outputs
for key, value in outputs.items():
    print(f"{key}: {value['value']}")

# Extrair o valor do output desejado
desired_output = outputs.get('your_output_name', {}).get('value', '')
print(f"Output desejado: {desired_output}")

# Escrever o valor no arquivo GITHUB_ENV
env_file = os.getenv('GITHUB_ENV')
with open(env_file, "a") as myfile:
    myfile.write(f"{output_name.upper()}={desired_output}\n")

