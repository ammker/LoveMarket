<template>
    <Header />
    <swiper
        :modules="modules"
        :slides-per-view="1" 
        navigation
        :autoplay="{ delay: 3000, disableOnInteraction: false }"
        :pagination="{ clickable: true }"
        @swiper="onSwiper"
        @slideChange="onSlideChange"
    >
        <swiper-slide><img src="../assets/11.webp" alt=""></swiper-slide>
        <swiper-slide><img src="../assets/641.webp" alt=""></swiper-slide>
        <swiper-slide><img src="../assets/123.jpg" alt=""></swiper-slide>
    </swiper>
</template>
  
<script>
    import Header from '@/components/Header.vue'
    import { Navigation, Pagination, Autoplay, A11y } from 'swiper/modules';
    import { Swiper, SwiperSlide } from 'swiper/vue';
    import 'swiper/css';
    import 'swiper/css/pagination';
    import 'swiper/css/navigation';
    import 'swiper/css/scrollbar';
  
    export default {
        components: {
            Swiper,
            SwiperSlide,
            Header
        },
        setup() {
            const onSwiper = (swiper) => {
                console.log(swiper);
            };
            const onSlideChange = () => {
                console.log('slide change');
            };
            return {
                onSwiper,
                onSlideChange,
                modules: [Navigation, Pagination, Autoplay, A11y],
            };
        },
    }
</script>

<style scoped>
    .home {
        font-family: Arial, sans-serif;
    }
    .swiper {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
    img{
        width:100%;
        height: 100%;
    }
</style>
