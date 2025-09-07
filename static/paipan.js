// 排盘页面JavaScript功能

class PaiPanApp {
    constructor() {
        this.currentUser = null;
        this.baziData = null;
        this.init();
    }

    init() {
        this.setupEventListeners();
        this.loadUserData();
    }

    setupEventListeners() {
        // 导航标签切换
        const tabs = document.querySelectorAll('.tab');
        tabs.forEach(tab => {
            tab.addEventListener('click', (e) => {
                tabs.forEach(t => t.classList.remove('active'));
                e.target.classList.add('active');
                this.switchTab(e.target.textContent);
            });
        });

        // 柱子悬停效果
        const pillars = document.querySelectorAll('.ganzhi-pair');
        pillars.forEach(pillar => {
            pillar.addEventListener('mouseenter', (e) => {
                this.showPillarDetails(e.currentTarget);
            });
            pillar.addEventListener('mouseleave', (e) => {
                this.hidePillarDetails(e.currentTarget);
            });
        });
    }

    // 加载用户数据
    loadUserData() {
        try {
            // 这里可以从URL参数或localStorage获取用户信息
            const urlParams = new URLSearchParams(window.location.search);
            const name = urlParams.get('name');
            const birthDate = urlParams.get('birthDate');
            const birthTime = urlParams.get('birthTime');

            console.log('URL参数:', { name, birthDate, birthTime });

            if (name && birthDate && birthTime) {
                this.fetchBaziData({ name, birthDate, birthTime });
            } else {
                // 显示示例数据
                console.log('未找到URL参数，显示示例数据');
                this.displaySampleData();
            }
        } catch (error) {
            console.error('加载用户数据时出错:', error);
            this.displaySampleData();
        }
    }

    // 获取八字数据
    async fetchBaziData(userData) {
        try {
            console.log('正在获取八字数据:', userData);
            
            const response = await fetch('/api/bazi', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: userData.name,
                    birthDate: userData.birthDate,
                    birthTime: userData.birthTime
                })
            });

            if (response.ok) {
                const baziData = await response.json();
                console.log('获取到八字数据:', baziData);
                this.baziData = baziData;
                this.updateUI(baziData);
            } else {
                console.error('获取八字数据失败:', response.status, response.statusText);
                this.displaySampleData();
            }
        } catch (error) {
            console.error('网络错误:', error);
            this.displaySampleData();
        }
    }

    // 更新UI显示
    updateUI(baziData) {
        // 更新基本信息
        if (baziData.birthDate && baziData.birthTime) {
            document.getElementById('solar-date').textContent = this.formatDateTime(baziData.birthDate, baziData.birthTime);
        }
        
        // 更新八字柱
        this.updateBaziPillars(baziData.bazi);
        
        // 更新详细信息
        this.updateDetailTable(baziData.bazi);
        
            // 更新用户名显示
        if (baziData.name) {
            document.title = `${baziData.name} - 八字排盘`;
        }
        
        // 隐藏加载指示器
        this.hideLoading();
    }
    
    // 隐藏加载指示器
    hideLoading() {
        const loading = document.getElementById('loading');
        if (loading) {
            loading.classList.add('loading-hidden');
        }
    }

    // 更新八字柱显示
    updateBaziPillars(baziColumns) {
        const ganzhi = document.querySelectorAll('.ganzhi-pair');
        const stars = document.querySelectorAll('.star-item');
        
        baziColumns.forEach((column, index) => {
            if (ganzhi[index]) {
                const gan = ganzhi[index].querySelector('.gan');
                const zhi = ganzhi[index].querySelector('.zhi');
                const ganWuxing = ganzhi[index].querySelector('.gan-wuxing');
                const zhiWuxing = ganzhi[index].querySelector('.zhi-wuxing');

                gan.textContent = column.gan;
                zhi.textContent = column.zhi;
                ganWuxing.textContent = column.ganWuXing;
                zhiWuxing.textContent = column.zhiWuXing;

                // 设置五行颜色
                gan.className = `gan ${this.getWuxingClass(column.ganWuXing)}`;
                zhi.className = `zhi ${this.getWuxingClass(column.zhiWuXing)}`;
            }

            if (stars[index] && column.zhuXing) {
                stars[index].textContent = column.zhuXing;
            }
        });
    }

    // 更新详细信息表格
    updateDetailTable(baziColumns) {
        const rows = document.querySelectorAll('.table-row');
        
        // 更新藏干
        const cangganRow = rows[0];
        const cangganCells = cangganRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (cangganCells[index] && column.cangGan) {
                cangganCells[index].innerHTML = column.cangGan.map(gan => 
                    `<span class="canggan ${this.getWuxingClass(this.getGanWuxing(gan))}">${gan}${this.getGanWuxing(gan)}</span>`
                ).join('');
            }
        });

        // 更新副星
        const fuxingRow = rows[1];
        const fuxingCells = fuxingRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (fuxingCells[index] && column.fuXing) {
                fuxingCells[index].textContent = column.fuXing.join('\n');
            }
        });

        // 更新纳音
        const nayinRow = rows[2];
        const nayinCells = nayinRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (nayinCells[index] && column.naYin) {
                nayinCells[index].textContent = column.naYin;
            }
        });

        // 更新星运
        const xingyunRow = rows[3];
        const xingyunCells = xingyunRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (xingyunCells[index] && column.xingYun) {
                xingyunCells[index].textContent = column.xingYun;
            }
        });

        // 更新自坐
        const zizuoRow = rows[4];
        const zizuoCells = zizuoRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (zizuoCells[index] && column.ziZuo) {
                zizuoCells[index].textContent = column.ziZuo;
            }
        });

        // 更新空亡
        const kongwangRow = rows[5];
        const kongwangCells = kongwangRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (kongwangCells[index]) {
                kongwangCells[index].textContent = column.kongWang ? '空亡' : '';
            }
        });

        // 更新神煞
        const shenshaRow = rows[6];
        const shenshaCells = shenshaRow.querySelectorAll('.row-data');
        baziColumns.forEach((column, index) => {
            if (shenshaCells[index] && column.shenSha) {
                shenshaCells[index].innerHTML = Object.keys(column.shenSha).map(key => 
                    `<div class="shensha-item">${key}</div>`
                ).join('');
            }
        });
    }

    // 获取五行对应的CSS类
    getWuxingClass(wuxing) {
        const wuxingMap = {
            '木': 'wood',
            '火': 'fire',
            '土': 'earth',
            '金': 'metal',
            '水': 'water'
        };
        return wuxingMap[wuxing] || '';
    }

    // 获取天干对应的五行
    getGanWuxing(gan) {
        const ganWuxingMap = {
            '甲': '木', '乙': '木',
            '丙': '火', '丁': '火',
            '戊': '土', '己': '土',
            '庚': '金', '辛': '金',
            '壬': '水', '癸': '水'
        };
        return ganWuxingMap[gan] || '';
    }

    // 格式化日期时间
    formatDateTime(date, time) {
        // 将 "2024-03-15" 和 "14:30" 格式化为 "2024年03月15日 14:30"
        const [year, month, day] = date.split('-');
        return `${year}年${month}月${day}日 ${time}`;
    }

    // 切换标签页
    switchTab(tabName) {
        console.log(`切换到标签页: ${tabName}`);
        // 这里可以实现不同标签页的内容切换
    }

    // 显示柱子详情
    showPillarDetails(pillarElement) {
        // 可以添加悬停时显示详细信息的功能
        pillarElement.style.transform = 'translateY(-5px)';
    }

    // 隐藏柱子详情
    hidePillarDetails(pillarElement) {
        pillarElement.style.transform = 'translateY(0)';
    }

    // 显示示例数据
    displaySampleData() {
        // 使用图片中的示例数据
        const sampleData = {
            name: "芦文超",
            birthDate: "1995-08-12",
            birthTime: "20:00",
            bazi: [
                {
                    gan: "乙", zhi: "亥",
                    ganWuXing: "木", zhiWuXing: "水",
                    zhuXing: "比肩",
                    cangGan: ["壬", "甲"],
                    fuXing: ["正印", "劫财"],
                    naYin: "山头火",
                    xingYun: "死",
                    ziZuo: "死",
                    kongWang: false,
                    shenSha: { "国印": "亥" }
                },
                {
                    gan: "甲", zhi: "申",
                    ganWuXing: "木", zhiWuXing: "金",
                    zhuXing: "劫财",
                    cangGan: ["庚", "壬", "戊"],
                    fuXing: ["正官", "正印", "正财"],
                    naYin: "泉中水",
                    xingYun: "胎",
                    ziZuo: "绝",
                    kongWang: false,
                    shenSha: { "天乙": "申", "红艳": "申" }
                },
                {
                    gan: "乙", zhi: "亥",
                    ganWuXing: "木", zhiWuXing: "水",
                    zhuXing: "元男",
                    cangGan: ["壬", "甲"],
                    fuXing: ["正印", "劫财"],
                    naYin: "山头火",
                    xingYun: "死",
                    ziZuo: "死",
                    kongWang: false,
                    shenSha: { "国印": "亥", "十灵日": "亥" }
                },
                {
                    gan: "丙", zhi: "戌",
                    ganWuXing: "火", zhiWuXing: "土",
                    zhuXing: "伤官",
                    cangGan: ["戊", "辛", "丁"],
                    fuXing: ["正财", "七杀", "食神"],
                    naYin: "屋上土",
                    xingYun: "墓",
                    ziZuo: "墓",
                    kongWang: false,
                    shenSha: { "天喜": "戌", "龙德": "戌" }
                }
            ]
        };

        this.updateUI(sampleData);
    }
}

// 页面加载完成后初始化应用
document.addEventListener('DOMContentLoaded', () => {
    new PaiPanApp();
});

// 导出到全局以便其他脚本使用
window.PaiPanApp = PaiPanApp;