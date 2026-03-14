<script setup lang="ts">
import { NDivider, NGradientText, NTag } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { get } from '@/api/system/about'
import { useAppStore } from '@/store'
const srcSvglogo = ref('/logo.png')
import srcGitee from '@/assets/about_image/gitee.png'
import srcGithub from '@/assets/about_image/github.png'
import srcDocker from '@/assets/about_image/docker.png'
import srcBilibili from '@/assets/about_image/bilibili.png'
import srcYoutube from '@/assets/about_image/youtube.png'
import srcQQGroupQR from '@/assets/about_image/qq_group_qr2.png'
import { RoundCardModal } from '@/components/common'

interface Version {
  versionName: string
  versionCode: number
}

const appStore = useAppStore()
const versionName = ref('')
const qqGroupQRShow = ref(false)
const frontVersion = import.meta.env.VITE_APP_VERSION || 'unknown'

onMounted(() => {
  get<Version>().then((res) => {
    if (res.code === 0)
      versionName.value = res.data.versionName
  })
})
</script>

<template>
  <div class="pt-5">
    <div class="flex flex-col items-center justify-center">
      <img :src="srcSvglogo" width="100" height="100" alt="">
      <div class="text-3xl font-semibold">
        {{ $t('common.appName') }}
      </div>

    </div>

    <NDivider style="margin:10px 0">
      •
    </NDivider>
    <div class="flex flex-col items-center justify-center text-base">
      <div>
        {{ $t('apps.about.author') }}<a href="https://github.com/ShuaiZha" target="_blank" class="link">Vox</a>
      </div>

    </div>
  </div>
</template>

<style>
.link{
    color:rgb(0, 89, 255)
}
</style>
