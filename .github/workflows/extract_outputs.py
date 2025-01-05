import sys
import json

# Obter o nome do arquivo dos argumentos de linha de comando
state_file = sys.argv[1] if len(sys.argv) > 1 else 'default.tfstate'

# Carregar o arquivo de estado do Terraform
with open(state_file, 'r') as file:
    state = json.load(file)

# Extrair os outputs
outputs = state.get('outputs', {})

# Imprimir os outputs
for key, value in outputs.items():
    print(f"{key}: {value['value']}")
