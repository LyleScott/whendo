set -ex

STACK_NAME="whendo"

root_dir=$(pwd)
artifacts_dir="${root_dir}/infra/_build"
sam_artifact_s3_bucket='lyle-stash-house'

# Prepare.
rm -rf "${artifacts_dir}"
mkdir -p "${artifacts_dir}"
aws s3 mb "s3://${sam_artifact_s3_bucket}"

# Build deployment artifacts.
for i in "when" "do"; do
  cd "src/${i}"
  GOOS=linux go build -o "${artifacts_dir}/${i}" main.go
  cd "${artifacts_dir}"
  zip "${i}.zip" "${i}"
  cd "${root_dir}"
done

# Update infra and code.
aws cloudformation package \
    --template-file infra/formation.yml \
    --output-template-file ${artifacts_dir}/formation.compile.yml \
    --s3-bucket "${sam_artifact_s3_bucket}"
aws cloudformation deploy \
    --template-file ${artifacts_dir}/formation.compile.yml \
    --stack-name "${STACK_NAME}" \
    --capabilities CAPABILITY_IAM
