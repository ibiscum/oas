npm install --global mock-to-openapi

openapi-generator-cli generate -g go -o circular -i ./yaml/circular.yaml --additional-properties=withGoMod=false,packageName=circular,isGoSubmodule=true,disallowAdditionalPropertiesIfNotPresent=false

openapi-generator-cli generate -g go -o callbacks -i ./yaml/callbacks.yaml --additional-properties=withGoMod=false,packageName=callbacks,isGoSubmodule=true,disallowAdditionalPropertiesIfNotPresent=false

openapi-generator-cli generate -g go -o complex_nesting -i ./yaml/complex_nesting.yaml --additional-properties=withGoMod=false,packageName=complex_nesting,isGoSubmodule=true,disallowAdditionalPropertiesIfNotPresent=false,generateAliasAsModel=true

openapi-generator-cli generate -g go -o polymorphism -i ./yaml/polymorphism.yaml --additional-properties=withGoMod=false,packageName=polymorphism,isGoSubmodule=true,disallowAdditionalPropertiesIfNotPresent=false,generateAliasAsModel=true

