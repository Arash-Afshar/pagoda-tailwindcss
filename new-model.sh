#! /bin/bash

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <ModelName>"
    exit 1
fi

MODELNAME="$1"
echo "1. creating ent/schema/${MODELNAME,,}.go"
cp ent/schema/modelname.go ent/schema/${MODELNAME,,}.go
sed -i "s/ModelName/${MODELNAME}/g" ent/schema/${MODELNAME,,}.go
if ! grep "${MODELNAME}s" ent/schema/user.go > /dev/null; then
    sed -i "s/\(return \[\]ent.Edge{\)/\1\n\t\tedge.To(\"${MODELNAME}s\", ${MODELNAME}.Type),/g" ent/schema/user.go
fi

echo "2. creating pkg/handlers/model_${MODELNAME,,}.go"
cp pkg/handlers/model_modelname.go pkg/handlers/model_${MODELNAME,,}.go
sed -i "s/ModelName/${MODELNAME}/g" pkg/handlers/model_${MODELNAME,,}.go
sed -i "s/modelname/${MODELNAME,,}/g" pkg/handlers/model_${MODELNAME,,}.go
sed -i "s/modelName/${MODELNAME,}/g" pkg/handlers/model_${MODELNAME,,}.go

echo "3. creating templates/pages/model-${MODELNAME,,}.gohtml"
cp templates/pages/model-modelname.gohtml templates/pages/model-${MODELNAME,,}.gohtml
sed -i "s/ModelName/${MODELNAME}/g" templates/pages/model-${MODELNAME,,}.gohtml
sed -i "s/modelname/${MODELNAME,,}/g" templates/pages/model-${MODELNAME,,}.gohtml

echo "4. creating templates/pages/model-${MODELNAME,,}-form.gohtml"
cp templates/pages/model-modelname-form.gohtml templates/pages/model-${MODELNAME,,}-form.gohtml
sed -i "s/ModelName/${MODELNAME}/g" templates/pages/model-${MODELNAME,,}-form.gohtml
sed -i "s/modelname/${MODELNAME,,}/g" templates/pages/model-${MODELNAME,,}-form.gohtml

echo "5. creating templates/pages/model-${MODELNAME,,}-list.gohtml"
cp templates/pages/model-modelname-list.gohtml templates/pages/model-${MODELNAME,,}-list.gohtml
sed -i "s/ModelName/${MODELNAME}/g" templates/pages/model-${MODELNAME,,}-list.gohtml
sed -i "s/modelname/${MODELNAME,,}/g" templates/pages/model-${MODELNAME,,}-list.gohtml


echo "6. Add these to template.go"
if ! grep "PageModel${MODELNAME}" templates/templates.go > /dev/null; then
    sed -i "s|\(Model Pages --------\)|\1\n\tPageModel${MODELNAME}            Page = \"model-${MODELNAME,,}\"\n\tPageModel${MODELNAME}List        Page = \"model-${MODELNAME,,}-list\"\n\tPageModel${MODELNAME}Form        Page = \"model-${MODELNAME,,}-form\"|g" templates/templates.go
fi

echo "7. Updating dashboard.go"
if ! grep "${MODELNAME}" pkg/handlers/dashboard.go > /dev/null; then
    sed -i "s|\(ent.Client\)|\1\n\t\t${MODELNAME,,} *Model${MODELNAME}|g" pkg/handlers/dashboard.go
    sed -i "s|\(ORM\)|\1\n\th.${MODELNAME,,} = NewModel${MODELNAME}(c)|g" pkg/handlers/dashboard.go
    sed -i "s|\(Sub-Routes --------\)|\1\n\n\tfor _, route := range h.${MODELNAME,,}.Routes().routeMaps {\n\t\tauth.Add(route.verb, route.path, route.handler).Name = route.name\n\t}\n\th.dashboardData.RouteMapMetas = append(h.dashboardData.RouteMapMetas, h.${MODELNAME,,}.Routes())|g" pkg/handlers/dashboard.go

fi


echo "8. Edit the code, then run 'make ent-gen'"
