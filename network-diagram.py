from diagrams import Diagram, Cluster
from diagrams.onprem.proxmox import ProxmoxVE
from diagrams.onprem.network import Internet
from diagrams.generic.network import VPN, Switch, Firewall, Router
from diagrams.generic.os import IOS, Windows, Ubuntu
from diagrams.custom import Custom
from diagrams.onprem.container import Docker
from diagrams.onprem.groupware import Nextcloud
from diagrams.oci.governance import Audit

with Diagram("Home Setup", show=False):
    internet = Internet('Internet')
    router = Router('ISP')

    with Cluster('TP-Link ER605'):
        vpn = VPN('VPN')
        firewall = Firewall('Firewall')

    with Cluster('Personal / Work'):
        mac = IOS('Mac Mini')
        gaming_pc = Windows('Gaming PC')

    with Cluster('Subnet - 192.168.0.*'):
        homelab = ProxmoxVE('homelab')
        
        with Cluster("VLAN 10"):
            jellyfin = Docker("JellyFin")

        with Cluster("VLAN 20"):
            torrent_machine = Docker("Torrent Machine") 

        with Cluster("VLAN 30"):

            with Cluster('Cluster'):
                control1 = Custom('Control 1', './resources/kubernetes-node.png')
                control2 = Custom('Control 2', './resources/kubernetes-node.png')
                control3 = Custom('Control 3', './resources/kubernetes-node.png')
                worker1 = Custom('Worker 1', './resources/kubernetes-node.png')
                worker2 = Custom('Worker 2', './resources/kubernetes-node.png')

                # nodes = control1 - control2 - control3 - worker1 - worker2
                nodes = worker1 - worker2 - control3 - control2 - control1

                nextcloud = Nextcloud('Nextcloud')
                joplin = Audit('Joplin')


        
    switch = Switch('Unmanaged Switch')

    # 1. Main backbone connection up to the firewall
    internet >> router >> switch

    switch >> vpn >> firewall
    
    # 2. Branch 1: Firewall straight to Proxmox homelab
    firewall >> homelab
    
    homelab >> jellyfin

    homelab >> torrent_machine

    homelab >> nodes

    nodes >> nextcloud

    nodes >> joplin
    
    switch >> [mac, gaming_pc]

