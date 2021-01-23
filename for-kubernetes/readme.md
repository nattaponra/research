Step 1: Install Docker
1. Update the package list with the command:
```
sudo apt-get update
```

kubernetes_1.-Update-Ubuntu-dependencies
2. Next, install Docker with the command:
```
sudo apt-get install docker.io
```

3. Repeat the process on each server that will act as a node.

4. Check the installation (and version) by entering the following:
```
docker ––version
```


Step 2: Start and Enable Docker
1. Set Docker to launch at boot by entering the following:
```
sudo systemctl enable docker
```


2. Verify Docker is running:
```
sudo systemctl status docker
```
To start Docker if it’s not running:
```
sudo systemctl start docker
```

3. Repeat on all the other nodes.