#码头工人
#Stalwart Mail 服务器可用作 Docker 映像，其中包括 JMAP、IMAP 和 SMTP 服务器。要开始使用，请拉取映像，例如：mail-server:latest
#
docker pull stalwartlabs/mail-server:latest
#
#然后，在主机上创建一个目录，您将在其中存储邮件服务器的配置文件和数据，例如：
#
mkdir /yushu/sofsware/mail
#
#完成设置说明后，启动 Stalwart Mail 服务器容器：
#
docker run -d -ti -p 11443:443 -p 11888:8080 \
             -p 25:25 -p 587:587 -p 465:465 \
             -p 143:143 -p 993:993 -p 4190:4190 \
             -p 110:110 -p 995:995 \
             -v /yushu/sofsware/mail:/opt/stalwart-mail \
             --name stalwart-mail stalwartlabs/mail-server:latest
#
#确保替换为您在上面创建的目录的路径。请注意，没有必要公开所有这些端口，请阅读入门文档以了解更多信息。<STALWART_DIR>
#
#登录到 Web 界面
#执行以获取系统的管理员帐户和密码：docker logs stalwart-mail
#
$ docker logs stalwart-mail
✅ Configuration file written to /opt/stalwart-mail/etc/config.toml
🔑 Your administrator account is 'admin' with password 'w95Yuiu36E'.
#
#利用此信息，您可以登录到 Web 界面 。http://yourserver.org:8080/login
#
#选择数据存储位置
#登录后，转到 > 部分并配置数据、blob、全文和查找存储。阅读 入门 小节 以了解有关可用选项的更多详细信息。SettingsStorage
#
#如果您想要外部身份验证目录（如 LDAP 或 SQL），请转到 > 部分并配置身份验证后端。SettingsAuthentication
#
#自选
#Stalwart Mail Server 预先配置为所有商店的默认后端。如果您对默认配置感到满意，可以跳过此步骤。RocksDB
#
#配置主机名和域
#接下来，确保 > > 中的服务器主机名正确无误。然后，在 > > 中添加您的主域名。创建域后，界面将显示您需要添加到域的 DNS 设置中的 DNS 记录。SettingsServerNetworkManagementDirectoryDomains
#
#例如：
#
#MX  example.org.                      10 mail.example.org.
#TXT 202404e._domainkey.example.org.   v=DKIM1; k=ed25519; h=sha256; p=MCowBQYDK2VwAyEAOT2JN9F8SLTVFNEODDu22SD9RJDC282mugCAeXkzjH0=
#TXT 202404r._domainkey.example.org.   v=DKIM1; k=rsa; h=sha256; p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAykeYJjv5N0AlnJ8gKF+/8qjbStiMFWvPg+p3JPh96GPXEN6l9W/Ee6Lag6i3vLyTVH5dnRVRBhfWhc+Dc0nKreZe4f5i4L5M4RI31+RpEgu4bCmncUIk2WzJgGBW5XbiOwXjge6OKWtJQN9d8Lc1AuryL5xeged9iS6xd/+EJz4WxAf18U+j38xmAm8fJUTBnQVeb/AZup+voSKAS59jyumsb0jQtXfX5xnwTFXdiX2OF8LRrmmNs/ObHozgHftxAv+YCiSU4bqSlKNPQIrN5kk1YnZDnLlc1Gr66AWlmdUVE7PWtZPTy4f8+uHO93EW3WUxLmynZm+Syn9FTJC2uwIDAQAB
#TXT mail.example.org.                 v=spf1 a -all ra=postmaster
#TXT example.org.                      v=spf1 mx -all ra=postmaster
#TXT _dmarc.example.org.               v=DMARC1; p=reject; rua=mailto:postmaster@example.org; ruf=mailto:postmaster@example.org
#
#
#提示
#根据您的设置，某些自动生成的记录可能是可选的，有关更多信息，请阅读了解 DNS 记录部分。
#
#启用 TLS
#Stalwart Mail Server 需要有效的 TLS 证书来保护服务器和客户端之间的连接。您可以通过以下方式之一启用 TLS：
#
#如果您的服务器已有 TLS 证书，则可以在 > > > 部分上传该证书。SettingsServerTLSCertificates
#如果您没有证书，则可以使用 ACME 从 Let's Encrypt 启用自动 TLS 证书。要启用 ACME，请转到 > > > 部分，并将 Let's Encrypt 添加为您的 ACME 提供商，确保您的服务器主机名列为主题名称之一。Stalwart 支持 和 质询，如果您不确定要使用哪一个，请阅读 ACME 质询类型文档。SettingsServerTLSACME Providerstls-alpn-01dns-01http-01
#如果您在 Traefik、Caddy、HAProxy 或 NGINX 等反向代理后面运行 Stalwart，则应跳过此步骤并在反向代理中配置 TLS。
#重启容器
#完成设置说明后，重新启动容器：
#
#$ docker restart stalwart-mail
#
#后续步骤
#如果您已选择使用内部目录，则现在可以在 > > 部分添加您的用户。如果您选择了外部目录，则需要在目录服务器中创建用户。ManagementDirectoryAccounts
#
#如果一切顺利，您的用户现在应该能够连接到服务器并发送和接收电子邮件。如果您无法连接到服务器，请从 web-admin 或 under 检查日志文件是否有任何错误。<INSTALL_DIR>/logs
#
#如果您有任何疑问，请查看 FAQ 部分或在社区论坛中发起讨论。
#
#提示
#在公开访问您的服务器之前，建议禁用任何未使用的服务以增强安全性。
#
#设置演示