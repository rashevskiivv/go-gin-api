mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
    var rootUser = '$MONGO_INITDB_ROOT_USERNAME';
    var rootPassword = '$MONGO_INITDB_ROOT_PASSWORD';
    var admin = db.getSiblingDB('admin');
    admin.auth(rootUser, rootPassword);

    var user = '$MONGO_INITDB_ROOT_USERNAME';
    var passwd = '$MONGO_INITDB_ROOT_PASSWORD';
    db.createUser({user: user, pwd: passwd, roles: ["readWrite"]});
EOF

#set -e
#
#mongosh <<EOF
#use $MONGO_INITDB_DATABASE
#
#db.createUser({
#  user: '$MONGO_INITDB_ROOT_USERNAME',
#  pwd: '$MONGO_INITDB_ROOT_PASSWORD',
#  roles: [{
#    role: 'readWrite',
#    db: '$MONGO_INITDB_DATABASE'
#  }]
#})
#EOF