package openai

var (
	Props = `{
    "assetPrefix": "",
    "buildId": "",
    "gssp": true,
    "isExperimentalCompile": false,
    "isFallback": false,
    "page": "/[[...default]]",
    "props": {
      "__N_SSP": true,
      "pageProps": {
        "serviceAnnouncement": {},
        "statsig": {
          "payload": {
            "feature_gates": {
              "458009956": { "name": "458009956", "value": true },
              "740954505": { "name": "740954505", "value": true }
            },
            "hash_used": "djb2",
            "layer_configs": {
              "2723963139": {
                "name": "2723963139",
                "value": {
                  "config": {
                    "impl": "rm_score",
                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                  },
                  "is_AG8PqS2q_enabled": true,
                  "is_conversation_model_switching_allowed": true,
                  "is_dynamic_model_enabled": true,
                  "show_message_model_info": true,
                  "show_message_regenerate_model_selector": true,
                  "show_message_regenerate_model_selector_on_every_message": true,
                  "show_rate_limit_downgrade_banner": true
                }
              },
              "3048336830": {
                "name": "3048336830",
                "value": { "is-enabled": true }
              }
            }
          },
          "user": { "userID": "user-xyhelper" }
        }
      }
    },
    "query": {},
    "scriptLoader": []
  }`
)
