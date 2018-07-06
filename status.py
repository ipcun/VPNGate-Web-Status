from softether.api import SoftEtherAPI

api = SoftEtherAPI('127.0.0.1', 443, 'password')

api.connect()
api.authenticate()

print(str(api.get_server_status()["NumSessionsTotal"]) + "  Users Connected")
print(str(api.get_server_status()["NumTcpConnections"]) + "  TCPConnections")

api.disconnect()