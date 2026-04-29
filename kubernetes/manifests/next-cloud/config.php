<?php
$CONFIG = array (
  'htaccess.RewriteBase' => '/',
  'memcache.local' => '\\OC\\Memcache\\APCu',
  'apps_paths' => 
  array (
    0 => 
    array (
      'path' => '/var/www/html/apps',
      'url' => '/apps',
      'writable' => false,
    ),
    1 => 
    array (
      'path' => '/var/www/html/custom_apps',
      'url' => '/custom_apps',
      'writable' => true,
    ),
  ),
  'upgrade.disable-web' => true,
  'instanceid' => 'oc7o1f4qinf7',
  'passwordsalt' => 'nC7E/SLyizucayeiKiFBayGsyf8Yrg',
  'secret' => 'CRPxZzDZWN2uoEIoeCOn/Hu8KWrVh+QxsGSgtMGOoTQYogMt',
  'trusted_domains' => 
  array (
    0 => '192.168.1.77:8080',
    1 => 'nextcloud.local',
    2 => '192.168.1.92:8080',
    3 => '192.168.1.92',
  ),
  'datadirectory' => '/var/www/html/data',
  'dbtype' => 'pgsql',
  'version' => '33.0.2.2',
  'overwrite.cli.url' => 'http://192.168.1.77:8080',
  'dbname' => 'next_cloud_db',
  'dbhost' => 'nextcloud-service',
  'dbtableprefix' => 'oc_',
  'dbuser' => 'oc_admin',
  'dbpassword' => 'F6uCeTtOXBVcwCLdMemz7CPiUAQGs6',
  'installed' => true,
);
